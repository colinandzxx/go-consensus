/*
 * Copyright (c) 2019
 *
 * This project is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * @File: bytes_test.go
 * @LastModified: 2019-10-08 17:32:12
 */

package types

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes_String(t *testing.T) {
	expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
	input := Bytes{
		1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
	}
	output := input.String()
	if output != expected {
		t.Errorf("Expected %s got %s", expected, output)
	}
}

func TestBytes_Format(t *testing.T) {
	{
		expected := "[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 21 22 23 24 25 26 27 28 29 210 211 212 213 214 215 216]"
		input := Bytes{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%v", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}

	{
		expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
		input := Bytes{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%x", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}
}

func TestByte32_String(t *testing.T) {
	expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
	input := Byte32{
		1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
	}
	output := input.String()
	if output != expected {
		t.Errorf("Expected %s got %s", expected, output)
	}
}

func TestByte32_Format(t *testing.T) {
	{
		expected := "[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 21 22 23 24 25 26 27 28 29 210 211 212 213 214 215 216]"
		input := Byte32{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%v", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}

	{
		expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
		input := Byte32{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%x", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}
}

func TestByte64_String(t *testing.T) {
	expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8" +
		"0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
	input := Byte64{
		1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
	}
	output := input.String()
	if output != expected {
		t.Errorf("Expected %s got %s", expected, output)
	}

	//t.Logf("String(): %s", b32.String())
	//t.Logf("fmt ori: %v", b32)
	//t.Logf("fmt hex: %x", b32)
}

func TestByte64_Format(t *testing.T) {
	{
		expected := "[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 21 22 23 24 25 26 27 28 29 210 211 212 213 214 215 216 " +
			"1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 21 22 23 24 25 26 27 28 29 210 211 212 213 214 215 216]"
		input := Byte64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
			1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%v", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}

	{
		expected := "0x0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8" +
			"0102030405060708090a0b0c0d0e0f1015161718191a1b1c1dd2d3d4d5d6d7d8"
		input := Byte64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
			1,  2,  3,  4,  5,  6,  7,  8,  9,  10,  11,  12,  13,  14,  15,  16,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 210, 211, 212, 213, 214, 215, 216,
		}
		output := bytes.Buffer{}
		fmt.Fprintf(&output, "%x", input)
		if output.String() != expected {
			t.Errorf("Expected %s got %s", expected, output.String())
		}
	}
}
