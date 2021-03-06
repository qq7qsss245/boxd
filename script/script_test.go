// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package script

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/BOXFoundation/boxd/core/pb"
	"github.com/BOXFoundation/boxd/core/types"
	"github.com/BOXFoundation/boxd/crypto"
	"github.com/facebookgo/ensure"
)

var (
	// TODO: to be wrapped in a helper function
	outPoint = types.OutPoint{
		Hash:  crypto.HashType{0x0010},
		Index: 0,
	}
	txIn = &types.TxIn{
		PrevOutPoint: outPoint,
		ScriptSig:    []byte{},
		Sequence:     0,
	}
	vIn = []*types.TxIn{
		txIn,
	}
	txOut = &corepb.TxOut{
		Value:        1,
		ScriptPubKey: []byte{},
	}
	vOut = []*corepb.TxOut{txOut}
	tx   = &types.Transaction{
		Version:  1,
		Vin:      vIn,
		Vout:     vOut,
		Magic:    1,
		LockTime: 0,
	}

	p2SHScriptBytes = []byte{
		byte(OPHASH160),
		0x14,                         // 160-bit redeemp script hash length: 20 bytes
		0x00, 0x01, 0x02, 0x03, 0x04, // 160-bit redeemp script hash: begining
		0x05, 0x06, 0x07, 0x08, 0x09,
		0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x13, // 160-bit redeemp script hash: end
		byte(OPEQUAL),
	}
)

// test script not dependent on a tx
func TestNonTxScriptEvaluation(t *testing.T) {
	script := NewScript().AddOpCode(OP8).AddOpCode(OP6).AddOpCode(OPADD).AddOpCode(OP14).AddOpCode(OPEQUAL)
	err := script.evaluate(nil, 0)
	ensure.Nil(t, err)
	script2 := NewScriptFromBytes(*script)
	ensure.DeepEqual(t, script2, script)

	script = NewScript().AddOpCode(OP8).AddOpCode(OP6).AddOpCode(OPADD).AddOpCode(OP11).AddOpCode(OPEQUAL)
	err = script.evaluate(nil, 0)
	ensure.NotNil(t, err)

	script = NewScript().AddOpCode(OP8).AddOpCode(OP6).AddOpCode(OPADD).AddOpCode(OP11).AddOpCode(OPEQUALVERIFY)
	err = script.evaluate(nil, 0)
	ensure.NotNil(t, err)

	script = NewScript().AddOpCode(OP8).AddOpCode(OP6).AddOpCode(OPSUB).AddOpCode(OP2).AddOpCode(OPEQUAL)
	err = script.evaluate(nil, 0)
	ensure.Nil(t, err)

	script = NewScript().AddOpCode(OP6).AddOpCode(OPDUP).AddOpCode(OPSUB).AddOpCode(OP0).AddOpCode(OPEQUAL)
	err = script.evaluate(nil, 0)
	ensure.Nil(t, err)

	script = NewScript().AddOpCode(OPDROP)
	err = script.evaluate(nil, 0)
	ensure.NotNil(t, err)

	script = NewScript().AddOpCode(OPTRUE).AddOpCode(OP16).AddOpCode(OPDROP)
	err = script.evaluate(nil, 0)
	ensure.Nil(t, err)

	script = NewScript().AddOpCode(OPFALSE).AddOpCode(OP16).AddOpCode(OPDROP)
	err = script.evaluate(nil, 0)
	ensure.NotNil(t, err)
}

func genP2PKHScript(appendOpDrop bool) (*Script, *Script, []byte) {
	// locking script: OPDUP, OPHASH160, testPubKeyHash, OPEQUALVERIFY, OPCHECKSIG
	scriptPubKey := NewScript().AddOpCode(OPDUP).AddOpCode(OPHASH160).AddOperand(testPubKeyHash).AddOpCode(OPEQUALVERIFY).AddOpCode(OPCHECKSIG)
	if appendOpDrop {
		scriptPubKey.AddOpCode(OP11).AddOpCode(OPDROP)
	}

	hash, _ := CalcTxHashForSig([]byte(*scriptPubKey), tx, 0)
	sig, _ := crypto.Sign(testPrivKey, hash)
	sigBytes := sig.Serialize()
	// unlocking script: sig, testPubKey
	scriptSig := NewScript().AddOperand(sigBytes).AddOperand(testPubKeyBytes)

	return scriptSig, scriptPubKey, sigBytes
}

// test p2pkh script
func TestP2PKH(t *testing.T) {
	scriptSig, scriptPubKey, _ := genP2PKHScript(false)
	err := Validate(scriptSig, scriptPubKey, tx, 0)
	ensure.Nil(t, err)
	ensure.DeepEqual(t, scriptSig.GetSigOpCount(), 0)
	ensure.DeepEqual(t, scriptPubKey.GetSigOpCount(), 1)

	// Append anything and immediately drop it to test OP_DROP; shall not affect script validity
	scriptSig, scriptPubKey, _ = genP2PKHScript(true)
	err = Validate(scriptSig, scriptPubKey, tx, 0)
	ensure.Nil(t, err)
}

func genP2SHScript() (*Script, *Script) {
	// redeem script
	redeemScript := NewScript().AddOperand(testPubKeyBytes).AddOpCode(OPCHECKSIG)
	redeemScriptHash := crypto.Hash160(*redeemScript)

	// locking script
	scriptPubKey := NewScript().AddOpCode(OPHASH160).AddOperand(redeemScriptHash).AddOpCode(OPEQUAL)

	// Note: use redeemScript, not scriptPubKey, because the former is checked against signature with OP_CHECKSIG
	hash, _ := CalcTxHashForSig([]byte(*redeemScript), tx, 0)
	sig, _ := crypto.Sign(testPrivKey, hash)
	sigBytes := sig.Serialize()
	// unlocking script: signature <redeemScript>
	// Note: <redeemScript> is serialized, i.e., AddOperand not AddScript
	scriptSig := NewScript().AddOperand(sigBytes).AddOperand(*redeemScript)

	return scriptSig, scriptPubKey
}

// test p2pkh script
func TestP2SH(t *testing.T) {
	scriptSig, scriptPubKey := genP2SHScript()
	err := Validate(scriptSig, scriptPubKey, tx, 0)
	ensure.Nil(t, err)
}

// minSigCount: minimal number of signatures required
// sigCount: number of signatures included in unlocking script
func genMultisigScript(minSigCount, sigCount int) (*Script, *Script) {
	testPrivKey1, testPubKey1, _ := crypto.NewKeyPair()
	testPubKeyBytes1 := testPubKey1.Serialize()

	testPrivKey2, testPubKey2, _ := crypto.NewKeyPair()
	testPubKeyBytes2 := testPubKey2.Serialize()

	// locking script: m <Public Key A> <Public Key B> <Public Key C> 3 CHECKMULTISIG
	opM := OpCode(int(OP1) + minSigCount - 1)
	scriptPubKey := NewScript().AddOpCode(opM).AddOperand(testPubKeyBytes).AddOperand(testPubKeyBytes1).
		AddOperand(testPubKeyBytes2).AddOpCode(OP3).AddOpCode(OPCHECKMULTISIG)

	hash, _ := CalcTxHashForSig([]byte(*scriptPubKey), tx, 0)

	sigs := make([][]byte, 0)

	sig, _ := crypto.Sign(testPrivKey, hash)
	sigs = append(sigs, sig.Serialize())

	sig, _ = crypto.Sign(testPrivKey1, hash)
	sigs = append(sigs, sig.Serialize())

	sig, _ = crypto.Sign(testPrivKey2, hash)
	sigs = append(sigs, sig.Serialize())

	// unlocking script: sigA, sigB
	scriptSig := NewScript()
	for i := 0; i < sigCount; i++ {
		scriptSig.AddOperand(sigs[i])
	}

	return scriptSig, scriptPubKey
}

// test multisig script
func TestMultisig(t *testing.T) {
	for minSigCount := 1; minSigCount <= 3; minSigCount++ {
		for sigCount := 1; sigCount <= 3; sigCount++ {
			scriptSig, scriptPubKey := genMultisigScript(minSigCount, sigCount)
			err := Validate(scriptSig, scriptPubKey, tx, 0)
			if sigCount < minSigCount {
				ensure.NotNil(t, err)
			} else {
				ensure.Nil(t, err)
			}
		}
	}
}

func TestDisasm(t *testing.T) {
	script := NewScript().AddOpCode(OP8).AddOpCode(OP6).AddOpCode(OPADD).AddOpCode(OP14).AddOpCode(OPEQUAL)
	ensure.DeepEqual(t, script.Disasm(), "OP_8 OP_6 OP_ADD OP_14 OP_EQUAL")

	// not enough data to push
	script.AddOpCode(OPPUSHDATA1)
	ensure.DeepEqual(t, script.Disasm(), "OP_8 OP_6 OP_ADD OP_14 OP_EQUAL [Error: OP_PUSHDATA1 has not enough data]")

	scriptSig, scriptPubKey, sigBytes := genP2PKHScript(false)
	expectedScriptStrs := []string{hex.EncodeToString(sigBytes), hex.EncodeToString(testPubKeyBytes), "OP_CODESEPARATOR",
		"OP_DUP", "OP_HASH160", hex.EncodeToString(testPubKeyHash), "OP_EQUALVERIFY", "OP_CHECKSIG"}
	catScript := NewScript().AddScript(scriptSig).AddOpCode(OPCODESEPARATOR).AddScript(scriptPubKey)
	ensure.DeepEqual(t, catScript.Disasm(), strings.Join(expectedScriptStrs, " "))
}

func TestIsPayToScriptHash(t *testing.T) {
	p2SHScript := NewScriptFromBytes(p2SHScriptBytes)
	ensure.True(t, p2SHScript.IsPayToScriptHash())

	p2SHScriptBytes[0] = byte(OPHASH256)
	p2SHScript = NewScriptFromBytes(p2SHScriptBytes)
	ensure.False(t, p2SHScript.IsPayToScriptHash())
	// recover
	p2SHScriptBytes[0] = byte(OPHASH160)

	p2SHScriptBytes[len(p2SHScriptBytes)-1] = byte(OPEQUALVERIFY)
	p2SHScript = NewScriptFromBytes(p2SHScriptBytes)
	ensure.False(t, p2SHScript.IsPayToScriptHash())
	// recover
	p2SHScriptBytes[len(p2SHScriptBytes)-1] = byte(OPEQUAL)

	p2SHScriptBytes[1] = 0x15
	p2SHScript = NewScriptFromBytes(p2SHScriptBytes)
	ensure.False(t, p2SHScript.IsPayToScriptHash())
	// recover
	p2SHScriptBytes[1] = 0x14

	p2SHScriptBytes = append(p2SHScriptBytes[:5], p2SHScriptBytes[6:]...)
	p2SHScript = NewScriptFromBytes(p2SHScriptBytes)
	ensure.False(t, p2SHScript.IsPayToScriptHash())
}

func TestIsPayToPubKeyHash(t *testing.T) {
	p2PKHScriptBytes := []byte{
		byte(OPDUP),
		byte(OPHASH160),
		0x14,                         // 160-bit public key hash length: 20 bytes
		0x00, 0x01, 0x02, 0x03, 0x04, // 160-bit public key hash: begining
		0x05, 0x06, 0x07, 0x08, 0x09,
		0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
		0x0F, 0x10, 0x11, 0x12, 0x13, // 160-bit public key hash: end
		byte(OPEQUALVERIFY),
		byte(OPCHECKSIG),
	}
	p2PKHScript := NewScriptFromBytes(p2PKHScriptBytes)
	ensure.True(t, p2PKHScript.IsPayToPubKeyHash())
}

func TestExtractAddress(t *testing.T) {
	// general tx
	_, scriptPubKey, _ := genP2PKHScript(false)
	addr, err := scriptPubKey.ExtractAddress()
	ensure.Nil(t, err)
	expectedAddr, _ := types.NewAddressFromPubKey(testPubKey)
	ensure.DeepEqual(t, expectedAddr, addr)

	_, scriptPubKey, _ = genP2PKHScript(true)
	_, err = scriptPubKey.ExtractAddress()
	ensure.NotNil(t, err)

	// p2sh
	_, scriptPubKey = genP2SHScript()
	_, err = scriptPubKey.ExtractAddress()
	ensure.NotNil(t, err)
}

func TestGetNthOp(t *testing.T) {
	// OPDUP, OPHASH160, testPubKeyHash, OPEQUALVERIFY, OPCHECKSIG
	_, scriptPubKey, _ := genP2PKHScript(false)

	// pc starts from 0
	opCode, _, _, _ := scriptPubKey.getNthOp(0 /* start pc */, 0 /* n-th */)
	ensure.DeepEqual(t, opCode, OPDUP)
	opCode, _, _, _ = scriptPubKey.getNthOp(0 /* start pc */, 1 /* n-th */)
	ensure.DeepEqual(t, opCode, OPHASH160)
	_, operand, _, _ := scriptPubKey.getNthOp(0 /* start pc */, 2 /* n-th */)
	ensure.DeepEqual(t, len(operand), 20)
	opCode, _, _, _ = scriptPubKey.getNthOp(0 /* start pc */, 3 /* n-th */)
	ensure.DeepEqual(t, opCode, OPEQUALVERIFY)
	opCode, _, _, _ = scriptPubKey.getNthOp(0 /* start pc */, 4 /* n-th */)
	ensure.DeepEqual(t, opCode, OPCHECKSIG)
	opCode, _, _, err := scriptPubKey.getNthOp(0 /* start pc */, 5 /* n-th */)
	ensure.NotNil(t, err)

	// moves pc
	opCode, _, pc, _ := scriptPubKey.getNthOp(0 /* start pc */, 0 /* n-th */)
	ensure.DeepEqual(t, opCode, OPDUP)
	opCode, _, pc, _ = scriptPubKey.getNthOp(pc /* start pc */, 0 /* n-th */)
	ensure.DeepEqual(t, opCode, OPHASH160)

	// pc stays
	_, operand, _, _ = scriptPubKey.getNthOp(pc /* start pc */, 0 /* n-th */)
	ensure.DeepEqual(t, len(operand), 20)
	opCode, _, _, _ = scriptPubKey.getNthOp(pc /* start pc */, 1 /* n-th */)
	ensure.DeepEqual(t, opCode, OPEQUALVERIFY)
	opCode, _, _, _ = scriptPubKey.getNthOp(pc /* start pc */, 2 /* n-th */)
	ensure.DeepEqual(t, opCode, OPCHECKSIG)
	opCode, _, _, err = scriptPubKey.getNthOp(pc /* start pc */, 3 /* n-th */)
	ensure.NotNil(t, err)
}
