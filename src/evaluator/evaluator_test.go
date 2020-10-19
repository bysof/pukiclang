package evaluator_test

import (
    "testing"

    "github.com/ythosa/pukiclang/src/lexer"
    "github.com/ythosa/pukiclang/src/object"
    "github.com/ythosa/pukiclang/src/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {"228", 228},
        {"1337", 1337},
    }

    for _, tt := range tests {

    }
}

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()

    return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
    result, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
        return false
    }

    if result.Value != expected {
        t.Errorf("object has wrong value. got=%d, want=%d",
            result.Value, expected)
    }

    return true
}