// Code generated from EventRule.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 318,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 75, 10, 6, 3, 7, 3, 7, 3, 7,
	5, 7, 80, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 24, 5, 24, 178, 10, 24, 3, 24, 6, 24, 181, 10, 24,
	13, 24, 14, 24, 182, 3, 24, 3, 24, 6, 24, 187, 10, 24, 13, 24, 14, 24,
	188, 5, 24, 191, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25,
	220, 10, 25, 3, 26, 3, 26, 3, 26, 7, 26, 225, 10, 26, 12, 26, 14, 26, 228,
	11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27, 234, 10, 27, 12, 27, 14, 27,
	237, 11, 27, 3, 27, 3, 27, 3, 27, 6, 27, 242, 10, 27, 13, 27, 14, 27, 243,
	3, 27, 6, 27, 247, 10, 27, 13, 27, 14, 27, 248, 5, 27, 251, 10, 27, 3,
	27, 3, 27, 6, 27, 255, 10, 27, 13, 27, 14, 27, 256, 5, 27, 259, 10, 27,
	5, 27, 261, 10, 27, 3, 27, 5, 27, 264, 10, 27, 3, 27, 3, 27, 3, 27, 7,
	27, 269, 10, 27, 12, 27, 14, 27, 272, 11, 27, 3, 27, 3, 27, 3, 27, 6, 27,
	277, 10, 27, 13, 27, 14, 27, 278, 3, 27, 6, 27, 282, 10, 27, 13, 27, 14,
	27, 283, 5, 27, 286, 10, 27, 3, 27, 3, 27, 6, 27, 290, 10, 27, 13, 27,
	14, 27, 291, 5, 27, 294, 10, 27, 5, 27, 296, 10, 27, 3, 27, 5, 27, 299,
	10, 27, 3, 27, 5, 27, 302, 10, 27, 7, 27, 304, 10, 27, 12, 27, 14, 27,
	307, 11, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5,
	29, 317, 10, 29, 3, 226, 2, 30, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51,
	27, 53, 28, 55, 29, 57, 2, 3, 2, 7, 3, 2, 47, 47, 3, 2, 50, 59, 5, 2, 67,
	92, 97, 97, 99, 124, 7, 2, 47, 47, 50, 59, 67, 92, 97, 97, 99, 124, 5,
	2, 11, 12, 15, 15, 34, 34, 2, 350, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3,
	2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37,
	3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2,
	45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2,
	2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 3, 59, 3, 2, 2, 2, 5, 61, 3, 2, 2,
	2, 7, 63, 3, 2, 2, 2, 9, 67, 3, 2, 2, 2, 11, 74, 3, 2, 2, 2, 13, 79, 3,
	2, 2, 2, 15, 81, 3, 2, 2, 2, 17, 84, 3, 2, 2, 2, 19, 86, 3, 2, 2, 2, 21,
	88, 3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 94, 3, 2, 2, 2, 27, 103, 3, 2,
	2, 2, 29, 116, 3, 2, 2, 2, 31, 119, 3, 2, 2, 2, 33, 126, 3, 2, 2, 2, 35,
	131, 3, 2, 2, 2, 37, 140, 3, 2, 2, 2, 39, 146, 3, 2, 2, 2, 41, 156, 3,
	2, 2, 2, 43, 163, 3, 2, 2, 2, 45, 174, 3, 2, 2, 2, 47, 177, 3, 2, 2, 2,
	49, 219, 3, 2, 2, 2, 51, 221, 3, 2, 2, 2, 53, 231, 3, 2, 2, 2, 55, 308,
	3, 2, 2, 2, 57, 316, 3, 2, 2, 2, 59, 60, 7, 42, 2, 2, 60, 4, 3, 2, 2, 2,
	61, 62, 7, 43, 2, 2, 62, 6, 3, 2, 2, 2, 63, 64, 7, 99, 2, 2, 64, 65, 7,
	112, 2, 2, 65, 66, 7, 102, 2, 2, 66, 8, 3, 2, 2, 2, 67, 68, 7, 113, 2,
	2, 68, 69, 7, 116, 2, 2, 69, 10, 3, 2, 2, 2, 70, 71, 7, 112, 2, 2, 71,
	72, 7, 113, 2, 2, 72, 75, 7, 118, 2, 2, 73, 75, 7, 35, 2, 2, 74, 70, 3,
	2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 12, 3, 2, 2, 2, 76, 80, 7, 63, 2, 2, 77,
	78, 7, 63, 2, 2, 78, 80, 7, 63, 2, 2, 79, 76, 3, 2, 2, 2, 79, 77, 3, 2,
	2, 2, 80, 14, 3, 2, 2, 2, 81, 82, 7, 35, 2, 2, 82, 83, 7, 63, 2, 2, 83,
	16, 3, 2, 2, 2, 84, 85, 7, 64, 2, 2, 85, 18, 3, 2, 2, 2, 86, 87, 7, 62,
	2, 2, 87, 20, 3, 2, 2, 2, 88, 89, 7, 64, 2, 2, 89, 90, 7, 63, 2, 2, 90,
	22, 3, 2, 2, 2, 91, 92, 7, 62, 2, 2, 92, 93, 7, 63, 2, 2, 93, 24, 3, 2,
	2, 2, 94, 95, 7, 101, 2, 2, 95, 96, 7, 113, 2, 2, 96, 97, 7, 112, 2, 2,
	97, 98, 7, 118, 2, 2, 98, 99, 7, 99, 2, 2, 99, 100, 7, 107, 2, 2, 100,
	101, 7, 112, 2, 2, 101, 102, 7, 117, 2, 2, 102, 26, 3, 2, 2, 2, 103, 104,
	7, 112, 2, 2, 104, 105, 7, 113, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107,
	7, 34, 2, 2, 107, 108, 7, 101, 2, 2, 108, 109, 7, 113, 2, 2, 109, 110,
	7, 112, 2, 2, 110, 111, 7, 118, 2, 2, 111, 112, 7, 99, 2, 2, 112, 113,
	7, 107, 2, 2, 113, 114, 7, 112, 2, 2, 114, 115, 7, 117, 2, 2, 115, 28,
	3, 2, 2, 2, 116, 117, 7, 107, 2, 2, 117, 118, 7, 112, 2, 2, 118, 30, 3,
	2, 2, 2, 119, 120, 7, 112, 2, 2, 120, 121, 7, 113, 2, 2, 121, 122, 7, 118,
	2, 2, 122, 123, 7, 34, 2, 2, 123, 124, 7, 107, 2, 2, 124, 125, 7, 112,
	2, 2, 125, 32, 3, 2, 2, 2, 126, 127, 7, 110, 2, 2, 127, 128, 7, 107, 2,
	2, 128, 129, 7, 109, 2, 2, 129, 130, 7, 103, 2, 2, 130, 34, 3, 2, 2, 2,
	131, 132, 7, 112, 2, 2, 132, 133, 7, 113, 2, 2, 133, 134, 7, 118, 2, 2,
	134, 135, 7, 34, 2, 2, 135, 136, 7, 110, 2, 2, 136, 137, 7, 107, 2, 2,
	137, 138, 7, 109, 2, 2, 138, 139, 7, 103, 2, 2, 139, 36, 3, 2, 2, 2, 140,
	141, 7, 116, 2, 2, 141, 142, 7, 103, 2, 2, 142, 143, 7, 105, 2, 2, 143,
	144, 7, 103, 2, 2, 144, 145, 7, 122, 2, 2, 145, 38, 3, 2, 2, 2, 146, 147,
	7, 112, 2, 2, 147, 148, 7, 113, 2, 2, 148, 149, 7, 118, 2, 2, 149, 150,
	7, 34, 2, 2, 150, 151, 7, 116, 2, 2, 151, 152, 7, 103, 2, 2, 152, 153,
	7, 105, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155, 7, 122, 2, 2, 155, 40,
	3, 2, 2, 2, 156, 157, 7, 103, 2, 2, 157, 158, 7, 122, 2, 2, 158, 159, 7,
	107, 2, 2, 159, 160, 7, 117, 2, 2, 160, 161, 7, 118, 2, 2, 161, 162, 7,
	117, 2, 2, 162, 42, 3, 2, 2, 2, 163, 164, 7, 112, 2, 2, 164, 165, 7, 113,
	2, 2, 165, 166, 7, 118, 2, 2, 166, 167, 7, 34, 2, 2, 167, 168, 7, 103,
	2, 2, 168, 169, 7, 122, 2, 2, 169, 170, 7, 107, 2, 2, 170, 171, 7, 117,
	2, 2, 171, 172, 7, 118, 2, 2, 172, 173, 7, 117, 2, 2, 173, 44, 3, 2, 2,
	2, 174, 175, 7, 46, 2, 2, 175, 46, 3, 2, 2, 2, 176, 178, 9, 2, 2, 2, 177,
	176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179, 181,
	9, 3, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 180, 3, 2,
	2, 2, 182, 183, 3, 2, 2, 2, 183, 190, 3, 2, 2, 2, 184, 186, 7, 48, 2, 2,
	185, 187, 9, 3, 2, 2, 186, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188,
	186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191, 3, 2, 2, 2, 190, 184,
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 48, 3, 2, 2, 2, 192, 193, 7, 86,
	2, 2, 193, 194, 7, 116, 2, 2, 194, 195, 7, 119, 2, 2, 195, 220, 7, 103,
	2, 2, 196, 197, 7, 86, 2, 2, 197, 198, 7, 84, 2, 2, 198, 199, 7, 87, 2,
	2, 199, 220, 7, 71, 2, 2, 200, 201, 7, 118, 2, 2, 201, 202, 7, 116, 2,
	2, 202, 203, 7, 119, 2, 2, 203, 220, 7, 103, 2, 2, 204, 205, 7, 72, 2,
	2, 205, 206, 7, 99, 2, 2, 206, 207, 7, 110, 2, 2, 207, 208, 7, 117, 2,
	2, 208, 220, 7, 103, 2, 2, 209, 210, 7, 72, 2, 2, 210, 211, 7, 67, 2, 2,
	211, 212, 7, 78, 2, 2, 212, 213, 7, 85, 2, 2, 213, 220, 7, 71, 2, 2, 214,
	215, 7, 104, 2, 2, 215, 216, 7, 99, 2, 2, 216, 217, 7, 110, 2, 2, 217,
	218, 7, 117, 2, 2, 218, 220, 7, 103, 2, 2, 219, 192, 3, 2, 2, 2, 219, 196,
	3, 2, 2, 2, 219, 200, 3, 2, 2, 2, 219, 204, 3, 2, 2, 2, 219, 209, 3, 2,
	2, 2, 219, 214, 3, 2, 2, 2, 220, 50, 3, 2, 2, 2, 221, 226, 7, 36, 2, 2,
	222, 225, 5, 57, 29, 2, 223, 225, 11, 2, 2, 2, 224, 222, 3, 2, 2, 2, 224,
	223, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 226, 224,
	3, 2, 2, 2, 227, 229, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 230, 7, 36,
	2, 2, 230, 52, 3, 2, 2, 2, 231, 235, 9, 4, 2, 2, 232, 234, 9, 5, 2, 2,
	233, 232, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235,
	236, 3, 2, 2, 2, 236, 263, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 260,
	7, 93, 2, 2, 239, 261, 7, 44, 2, 2, 240, 242, 9, 3, 2, 2, 241, 240, 3,
	2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2,
	2, 244, 261, 3, 2, 2, 2, 245, 247, 9, 3, 2, 2, 246, 245, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 251,
	3, 2, 2, 2, 250, 246, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2,
	2, 2, 252, 258, 7, 60, 2, 2, 253, 255, 9, 3, 2, 2, 254, 253, 3, 2, 2, 2,
	255, 256, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	259, 3, 2, 2, 2, 258, 254, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 261,
	3, 2, 2, 2, 260, 239, 3, 2, 2, 2, 260, 241, 3, 2, 2, 2, 260, 250, 3, 2,
	2, 2, 261, 262, 3, 2, 2, 2, 262, 264, 7, 95, 2, 2, 263, 238, 3, 2, 2, 2,
	263, 264, 3, 2, 2, 2, 264, 305, 3, 2, 2, 2, 265, 266, 7, 48, 2, 2, 266,
	270, 9, 4, 2, 2, 267, 269, 9, 5, 2, 2, 268, 267, 3, 2, 2, 2, 269, 272,
	3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 298, 3, 2,
	2, 2, 272, 270, 3, 2, 2, 2, 273, 295, 7, 93, 2, 2, 274, 296, 7, 44, 2,
	2, 275, 277, 9, 3, 2, 2, 276, 275, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278,
	276, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 296, 3, 2, 2, 2, 280, 282,
	9, 3, 2, 2, 281, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 281, 3, 2,
	2, 2, 283, 284, 3, 2, 2, 2, 284, 286, 3, 2, 2, 2, 285, 281, 3, 2, 2, 2,
	285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 293, 7, 60, 2, 2, 288,
	290, 9, 3, 2, 2, 289, 288, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 289,
	3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 294, 3, 2, 2, 2, 293, 289, 3, 2,
	2, 2, 293, 294, 3, 2, 2, 2, 294, 296, 3, 2, 2, 2, 295, 274, 3, 2, 2, 2,
	295, 276, 3, 2, 2, 2, 295, 285, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297,
	299, 7, 95, 2, 2, 298, 273, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 301,
	3, 2, 2, 2, 300, 302, 7, 48, 2, 2, 301, 300, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 304, 3, 2, 2, 2, 303, 265, 3, 2, 2, 2, 304, 307, 3, 2, 2, 2,
	305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 54, 3, 2, 2, 2, 307, 305,
	3, 2, 2, 2, 308, 309, 9, 6, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 8, 28,
	2, 2, 311, 56, 3, 2, 2, 2, 312, 313, 7, 94, 2, 2, 313, 317, 7, 36, 2, 2,
	314, 315, 7, 94, 2, 2, 315, 317, 7, 94, 2, 2, 316, 312, 3, 2, 2, 2, 316,
	314, 3, 2, 2, 2, 317, 58, 3, 2, 2, 2, 31, 2, 74, 79, 177, 182, 188, 190,
	219, 224, 226, 235, 243, 248, 250, 256, 258, 260, 263, 270, 278, 283, 285,
	291, 293, 295, 298, 301, 305, 316, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'and'", "'or'", "", "", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'contains'", "'not contains'", "'in'", "'not in'", "'like'", "'not like'",
	"'regex'", "'not regex'", "'exists'", "'not exists'", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "AND", "OR", "NOT", "EQU", "NEQ", "GT", "LT", "GTE", "LTE",
	"CONTAINS", "NOTCONTAINS", "IN", "NOTIN", "LIKE", "NOTLIKE", "REGEX", "NOTREGEX",
	"EXISTS", "NOTEXISTS", "COMMA", "NUMBER", "BOOLEAN", "STRING", "VAR", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "AND", "OR", "NOT", "EQU", "NEQ", "GT", "LT", "GTE", "LTE",
	"CONTAINS", "NOTCONTAINS", "IN", "NOTIN", "LIKE", "NOTLIKE", "REGEX", "NOTREGEX",
	"EXISTS", "NOTEXISTS", "COMMA", "NUMBER", "BOOLEAN", "STRING", "VAR", "WHITESPACE",
	"ESC",
}

type EventRuleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewEventRuleLexer(input antlr.CharStream) *EventRuleLexer {

	l := new(EventRuleLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "EventRule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// EventRuleLexer tokens.
const (
	EventRuleLexerT__0        = 1
	EventRuleLexerT__1        = 2
	EventRuleLexerAND         = 3
	EventRuleLexerOR          = 4
	EventRuleLexerNOT         = 5
	EventRuleLexerEQU         = 6
	EventRuleLexerNEQ         = 7
	EventRuleLexerGT          = 8
	EventRuleLexerLT          = 9
	EventRuleLexerGTE         = 10
	EventRuleLexerLTE         = 11
	EventRuleLexerCONTAINS    = 12
	EventRuleLexerNOTCONTAINS = 13
	EventRuleLexerIN          = 14
	EventRuleLexerNOTIN       = 15
	EventRuleLexerLIKE        = 16
	EventRuleLexerNOTLIKE     = 17
	EventRuleLexerREGEX       = 18
	EventRuleLexerNOTREGEX    = 19
	EventRuleLexerEXISTS      = 20
	EventRuleLexerNOTEXISTS   = 21
	EventRuleLexerCOMMA       = 22
	EventRuleLexerNUMBER      = 23
	EventRuleLexerBOOLEAN     = 24
	EventRuleLexerSTRING      = 25
	EventRuleLexerVAR         = 26
	EventRuleLexerWHITESPACE  = 27
)
