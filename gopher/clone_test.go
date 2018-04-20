package main

import (
	"github.com/yuin/gopher-lua"
	"testing"
)

func makeSource(L *lua.LState) {
	tbl := L.NewTable()
	L.SetTable(tbl, lua.LString("alpha"), lua.LString("beta"))
	L.SetGlobal("gamma", tbl)
}

func testDestinate(t *testing.T, L *lua.LState) {
	tbl := L.GetGlobal("gamma")
	if tbl.Type() != lua.LTTable {
		t.Fatal("Failed to copy table.")
		return
	}
	val := L.GetField(tbl, "alpha")
	if val != lua.LString("beta") {
		t.Fatalf("Failed to copy all of instance(%s)", val.String())
		return
	}
}

func TestClone(t *testing.T) {
	L1 := lua.NewState()

	makeSource(L1)
	L2 := Clone(L1)
	L1.Close()

	if L2 == nil {
		t.Fatal("Failed to create instance")
		return
	}
	testDestinate(t, L2)
	L2.Close()
}
