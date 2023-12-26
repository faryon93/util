package util

import "fmt"

// util
// Copyright (C) 2023 Maximilian Pachl

// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------
//  types
// ---------------------------------------------------------------------------------------

// FlagArray is a type which could be used with the flag package, to obtain
// a list from multiple flags with the same name but different values.
// e.g.: ./test --foo=bar --foo=test --foo=x
type FlagArray []string

// ---------------------------------------------------------------------------------------
//	public members
// ---------------------------------------------------------------------------------------

func (f *FlagArray) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

func (f *FlagArray) Set(value string) error {
	*f = append(*f, value)
	return nil
}
