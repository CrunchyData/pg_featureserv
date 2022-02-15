// Generated from CQL.g4 by ANTLR 4.7.

package cql // CQL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 92, 353,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 96, 10, 3, 12, 3, 14, 3,
	99, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 107, 10, 4, 12, 4,
	14, 4, 110, 11, 4, 3, 5, 5, 5, 113, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 122, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 129,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 137, 10, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 5, 10, 144, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 5, 11, 154, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 162, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18,
	3, 18, 5, 18, 181, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 191, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 205, 10, 22, 12, 22, 14,
	22, 208, 11, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 7, 24, 219, 10, 24, 12, 24, 14, 24, 222, 11, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 231, 10, 25, 12, 25, 14, 25, 234,
	11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 243, 10,
	26, 12, 26, 14, 26, 246, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 7, 27, 255, 10, 27, 12, 27, 14, 27, 258, 11, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 267, 10, 28, 12, 28, 14, 28,
	270, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 283, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 290, 10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 5, 30, 297, 10,
	30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 41, 3, 41, 5, 41, 323, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43,
	5, 43, 329, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43, 336, 10,
	43, 12, 43, 14, 43, 339, 11, 43, 3, 43, 3, 43, 3, 43, 7, 43, 344, 10, 43,
	12, 43, 14, 43, 347, 11, 43, 5, 43, 349, 10, 43, 3, 43, 3, 43, 3, 43, 2,
	4, 4, 6, 44, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 2, 4, 3, 2, 14, 15, 4, 2, 3, 3, 24, 24,
	2, 346, 2, 86, 3, 2, 2, 2, 4, 89, 3, 2, 2, 2, 6, 100, 3, 2, 2, 2, 8, 112,
	3, 2, 2, 2, 10, 121, 3, 2, 2, 2, 12, 128, 3, 2, 2, 2, 14, 130, 3, 2, 2,
	2, 16, 134, 3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 161,
	3, 2, 2, 2, 24, 163, 3, 2, 2, 2, 26, 165, 3, 2, 2, 2, 28, 167, 3, 2, 2,
	2, 30, 169, 3, 2, 2, 2, 32, 171, 3, 2, 2, 2, 34, 180, 3, 2, 2, 2, 36, 190,
	3, 2, 2, 2, 38, 192, 3, 2, 2, 2, 40, 197, 3, 2, 2, 2, 42, 200, 3, 2, 2,
	2, 44, 211, 3, 2, 2, 2, 46, 214, 3, 2, 2, 2, 48, 225, 3, 2, 2, 2, 50, 237,
	3, 2, 2, 2, 52, 249, 3, 2, 2, 2, 54, 261, 3, 2, 2, 2, 56, 273, 3, 2, 2,
	2, 58, 293, 3, 2, 2, 2, 60, 298, 3, 2, 2, 2, 62, 300, 3, 2, 2, 2, 64, 302,
	3, 2, 2, 2, 66, 304, 3, 2, 2, 2, 68, 306, 3, 2, 2, 2, 70, 308, 3, 2, 2,
	2, 72, 310, 3, 2, 2, 2, 74, 312, 3, 2, 2, 2, 76, 314, 3, 2, 2, 2, 78, 316,
	3, 2, 2, 2, 80, 322, 3, 2, 2, 2, 82, 324, 3, 2, 2, 2, 84, 326, 3, 2, 2,
	2, 86, 87, 5, 4, 3, 2, 87, 88, 7, 2, 2, 3, 88, 3, 3, 2, 2, 2, 89, 90, 8,
	3, 1, 2, 90, 91, 5, 6, 4, 2, 91, 97, 3, 2, 2, 2, 92, 93, 12, 3, 2, 2, 93,
	94, 7, 12, 2, 2, 94, 96, 5, 6, 4, 2, 95, 92, 3, 2, 2, 2, 96, 99, 3, 2,
	2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 5, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 100, 101, 8, 4, 1, 2, 101, 102, 5, 8, 5, 2, 102, 108, 3, 2,
	2, 2, 103, 104, 12, 3, 2, 2, 104, 105, 7, 11, 2, 2, 105, 107, 5, 8, 5,
	2, 106, 103, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 7, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 113, 7,
	13, 2, 2, 112, 111, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 3, 2, 2,
	2, 114, 115, 5, 10, 6, 2, 115, 9, 3, 2, 2, 2, 116, 122, 5, 12, 7, 2, 117,
	118, 7, 51, 2, 2, 118, 119, 5, 4, 3, 2, 119, 120, 7, 52, 2, 2, 120, 122,
	3, 2, 2, 2, 121, 116, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 122, 11, 3, 2,
	2, 2, 123, 129, 5, 14, 8, 2, 124, 129, 5, 16, 9, 2, 125, 129, 5, 18, 10,
	2, 126, 129, 5, 20, 11, 2, 127, 129, 5, 84, 43, 2, 128, 123, 3, 2, 2, 2,
	128, 124, 3, 2, 2, 2, 128, 125, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	127, 3, 2, 2, 2, 129, 13, 3, 2, 2, 2, 130, 131, 5, 22, 12, 2, 131, 132,
	7, 3, 2, 2, 132, 133, 5, 22, 12, 2, 133, 15, 3, 2, 2, 2, 134, 136, 5, 24,
	13, 2, 135, 137, 7, 13, 2, 2, 136, 135, 3, 2, 2, 2, 136, 137, 3, 2, 2,
	2, 137, 138, 3, 2, 2, 2, 138, 139, 9, 2, 2, 2, 139, 140, 5, 26, 14, 2,
	140, 17, 3, 2, 2, 2, 141, 143, 5, 24, 13, 2, 142, 144, 7, 13, 2, 2, 143,
	142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 146,
	7, 16, 2, 2, 146, 147, 5, 22, 12, 2, 147, 148, 7, 11, 2, 2, 148, 149, 5,
	22, 12, 2, 149, 19, 3, 2, 2, 2, 150, 151, 5, 24, 13, 2, 151, 153, 7, 17,
	2, 2, 152, 154, 7, 13, 2, 2, 153, 152, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2,
	154, 155, 3, 2, 2, 2, 155, 156, 7, 18, 2, 2, 156, 21, 3, 2, 2, 2, 157,
	162, 5, 24, 13, 2, 158, 162, 5, 26, 14, 2, 159, 162, 5, 28, 15, 2, 160,
	162, 5, 30, 16, 2, 161, 157, 3, 2, 2, 2, 161, 158, 3, 2, 2, 2, 161, 159,
	3, 2, 2, 2, 161, 160, 3, 2, 2, 2, 162, 23, 3, 2, 2, 2, 163, 164, 7, 39,
	2, 2, 164, 25, 3, 2, 2, 2, 165, 166, 7, 91, 2, 2, 166, 27, 3, 2, 2, 2,
	167, 168, 7, 38, 2, 2, 168, 29, 3, 2, 2, 2, 169, 170, 7, 10, 2, 2, 170,
	31, 3, 2, 2, 2, 171, 172, 7, 23, 2, 2, 172, 173, 7, 51, 2, 2, 173, 174,
	5, 34, 18, 2, 174, 175, 7, 57, 2, 2, 175, 176, 5, 34, 18, 2, 176, 177,
	7, 52, 2, 2, 177, 33, 3, 2, 2, 2, 178, 181, 5, 24, 13, 2, 179, 181, 5,
	36, 19, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 35, 3, 2, 2,
	2, 182, 191, 5, 38, 20, 2, 183, 191, 5, 40, 21, 2, 184, 191, 5, 44, 23,
	2, 185, 191, 5, 48, 25, 2, 186, 191, 5, 50, 26, 2, 187, 191, 5, 52, 27,
	2, 188, 191, 5, 54, 28, 2, 189, 191, 5, 56, 29, 2, 190, 182, 3, 2, 2, 2,
	190, 183, 3, 2, 2, 2, 190, 184, 3, 2, 2, 2, 190, 185, 3, 2, 2, 2, 190,
	186, 3, 2, 2, 2, 190, 187, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 189,
	3, 2, 2, 2, 191, 37, 3, 2, 2, 2, 192, 193, 7, 30, 2, 2, 193, 194, 7, 51,
	2, 2, 194, 195, 5, 58, 30, 2, 195, 196, 7, 52, 2, 2, 196, 39, 3, 2, 2,
	2, 197, 198, 7, 31, 2, 2, 198, 199, 5, 42, 22, 2, 199, 41, 3, 2, 2, 2,
	200, 201, 7, 51, 2, 2, 201, 206, 5, 58, 30, 2, 202, 203, 7, 57, 2, 2, 203,
	205, 5, 58, 30, 2, 204, 202, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204,
	3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 3, 2, 2, 2, 208, 206, 3, 2,
	2, 2, 209, 210, 7, 52, 2, 2, 210, 43, 3, 2, 2, 2, 211, 212, 7, 32, 2, 2,
	212, 213, 5, 46, 24, 2, 213, 45, 3, 2, 2, 2, 214, 215, 7, 51, 2, 2, 215,
	220, 5, 42, 22, 2, 216, 217, 7, 57, 2, 2, 217, 219, 5, 42, 22, 2, 218,
	216, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 220, 221,
	3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 224, 7, 52,
	2, 2, 224, 47, 3, 2, 2, 2, 225, 226, 7, 33, 2, 2, 226, 227, 7, 51, 2, 2,
	227, 232, 5, 58, 30, 2, 228, 229, 7, 57, 2, 2, 229, 231, 5, 58, 30, 2,
	230, 228, 3, 2, 2, 2, 231, 234, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232,
	233, 3, 2, 2, 2, 233, 235, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 235, 236,
	7, 52, 2, 2, 236, 49, 3, 2, 2, 2, 237, 238, 7, 34, 2, 2, 238, 239, 7, 51,
	2, 2, 239, 244, 5, 42, 22, 2, 240, 241, 7, 57, 2, 2, 241, 243, 5, 42, 22,
	2, 242, 240, 3, 2, 2, 2, 243, 246, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 244,
	245, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 247, 248,
	7, 52, 2, 2, 248, 51, 3, 2, 2, 2, 249, 250, 7, 35, 2, 2, 250, 251, 7, 51,
	2, 2, 251, 256, 5, 46, 24, 2, 252, 253, 7, 57, 2, 2, 253, 255, 5, 46, 24,
	2, 254, 252, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256,
	257, 3, 2, 2, 2, 257, 259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 260,
	7, 52, 2, 2, 260, 53, 3, 2, 2, 2, 261, 262, 7, 36, 2, 2, 262, 263, 7, 51,
	2, 2, 263, 268, 5, 36, 19, 2, 264, 265, 7, 57, 2, 2, 265, 267, 5, 36, 19,
	2, 266, 264, 3, 2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268,
	269, 3, 2, 2, 2, 269, 271, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 272,
	7, 52, 2, 2, 272, 55, 3, 2, 2, 2, 273, 274, 7, 37, 2, 2, 274, 275, 7, 51,
	2, 2, 275, 276, 5, 66, 34, 2, 276, 277, 7, 57, 2, 2, 277, 278, 5, 72, 37,
	2, 278, 282, 7, 57, 2, 2, 279, 280, 5, 74, 38, 2, 280, 281, 7, 57, 2, 2,
	281, 283, 3, 2, 2, 2, 282, 279, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283,
	284, 3, 2, 2, 2, 284, 285, 5, 68, 35, 2, 285, 286, 7, 57, 2, 2, 286, 289,
	5, 70, 36, 2, 287, 288, 7, 57, 2, 2, 288, 290, 5, 76, 39, 2, 289, 287,
	3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 7, 52,
	2, 2, 292, 57, 3, 2, 2, 2, 293, 294, 5, 60, 31, 2, 294, 296, 5, 62, 32,
	2, 295, 297, 5, 64, 33, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2,
	297, 59, 3, 2, 2, 2, 298, 299, 7, 38, 2, 2, 299, 61, 3, 2, 2, 2, 300, 301,
	7, 38, 2, 2, 301, 63, 3, 2, 2, 2, 302, 303, 7, 38, 2, 2, 303, 65, 3, 2,
	2, 2, 304, 305, 7, 38, 2, 2, 305, 67, 3, 2, 2, 2, 306, 307, 7, 38, 2, 2,
	307, 69, 3, 2, 2, 2, 308, 309, 7, 38, 2, 2, 309, 71, 3, 2, 2, 2, 310, 311,
	7, 38, 2, 2, 311, 73, 3, 2, 2, 2, 312, 313, 7, 38, 2, 2, 313, 75, 3, 2,
	2, 2, 314, 315, 7, 38, 2, 2, 315, 77, 3, 2, 2, 2, 316, 317, 5, 80, 41,
	2, 317, 318, 9, 3, 2, 2, 318, 319, 5, 80, 41, 2, 319, 79, 3, 2, 2, 2, 320,
	323, 5, 24, 13, 2, 321, 323, 5, 82, 42, 2, 322, 320, 3, 2, 2, 2, 322, 321,
	3, 2, 2, 2, 323, 81, 3, 2, 2, 2, 324, 325, 7, 76, 2, 2, 325, 83, 3, 2,
	2, 2, 326, 328, 5, 24, 13, 2, 327, 329, 7, 13, 2, 2, 328, 327, 3, 2, 2,
	2, 328, 329, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 7, 29, 2, 2, 331,
	348, 7, 51, 2, 2, 332, 337, 5, 26, 14, 2, 333, 334, 7, 57, 2, 2, 334, 336,
	5, 26, 14, 2, 335, 333, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2, 337, 335, 3,
	2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 349, 3, 2, 2, 2, 339, 337, 3, 2, 2,
	2, 340, 345, 5, 28, 15, 2, 341, 342, 7, 57, 2, 2, 342, 344, 5, 28, 15,
	2, 343, 341, 3, 2, 2, 2, 344, 347, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345,
	346, 3, 2, 2, 2, 346, 349, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 332,
	3, 2, 2, 2, 348, 340, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 7, 52,
	2, 2, 351, 85, 3, 2, 2, 2, 27, 97, 108, 112, 121, 128, 136, 143, 153, 161,
	180, 190, 206, 220, 232, 244, 256, 268, 282, 289, 296, 322, 328, 337, 345,
	348,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "'#'", "'$'", "'_'", "'\"'", "'%'", "'&'",
	"", "'('", "')'", "'['", "']'", "'*'", "'+'", "','", "'-'", "'.'", "'/'",
	"':'", "';'", "'?'", "'|'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "''''",
}
var symbolicNames = []string{
	"", "ComparisonOperator", "LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral",
	"AND", "OR", "NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "WILDCARD",
	"SINGLECHAR", "ESCAPECHAR", "NOCASE", "SpatialOperator", "TemporalOperator",
	"ArrayOperator", "EXISTS", "EXIST", "DOES", "IN", "POINT", "LINESTRING",
	"POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON", "GEOMETRYCOLLECTION",
	"ENVELOPE", "NumericLiteral", "Identifier", "IdentifierStart", "IdentifierPart",
	"ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR", "UNDERSCORE", "DOUBLEQUOTE", "PERCENT",
	"AMPERSAND", "QUOTE", "LEFTPAREN", "RIGHTPAREN", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET",
	"ASTERISK", "PLUS", "COMMA", "MINUS", "PERIOD", "SOLIDUS", "COLON", "SEMICOLON",
	"QUESTIONMARK", "VERTICALBAR", "BIT", "HEXIT", "UnsignedNumericLiteral",
	"SignedNumericLiteral", "ExactNumericLiteral", "ApproximateNumericLiteral",
	"Mantissa", "Exponent", "SignedInteger", "UnsignedInteger", "Sign", "TemporalLiteral",
	"Instant", "Interval", "InstantInInterval", "FullDate", "DateYear", "DateMonth",
	"DateDay", "UtcTime", "TimeZoneOffset", "TimeHour", "TimeMinute", "TimeSecond",
	"NOW", "WS", "CharacterStringLiteral", "QuotedQuote",
}

var ruleNames = []string{
	"cqlFilter", "booleanValueExpression", "booleanTerm", "booleanFactor",
	"booleanPrimary", "predicate", "binaryComparisonPredicate", "likePredicate",
	"betweenPredicate", "isNullPredicate", "scalarExpression", "propertyName",
	"characterLiteral", "numericLiteral", "booleanLiteral", "spatialPredicate",
	"geomExpression", "geomLiteral", "point", "linestring", "linestringDef",
	"polygon", "polygonDef", "multiPoint", "multiLinestring", "multiPolygon",
	"geometryCollection", "envelope", "coordinate", "xCoord", "yCoord", "zCoord",
	"westBoundLon", "eastBoundLon", "northBoundLat", "southBoundLat", "minElev",
	"maxElev", "temporalPredicate", "temporalExpression", "temporalLiteral",
	"inPredicate",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CQL struct {
	*antlr.BaseParser
}

func NewCQL(input antlr.TokenStream) *CQL {
	this := new(CQL)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CQL.g4"

	return this
}

// CQL tokens.
const (
	CQLEOF                       = antlr.TokenEOF
	CQLComparisonOperator        = 1
	CQLLT                        = 2
	CQLEQ                        = 3
	CQLGT                        = 4
	CQLNEQ                       = 5
	CQLGTEQ                      = 6
	CQLLTEQ                      = 7
	CQLBooleanLiteral            = 8
	CQLAND                       = 9
	CQLOR                        = 10
	CQLNOT                       = 11
	CQLLIKE                      = 12
	CQLILIKE                     = 13
	CQLBETWEEN                   = 14
	CQLIS                        = 15
	CQLNULL                      = 16
	CQLWILDCARD                  = 17
	CQLSINGLECHAR                = 18
	CQLESCAPECHAR                = 19
	CQLNOCASE                    = 20
	CQLSpatialOperator           = 21
	CQLTemporalOperator          = 22
	CQLArrayOperator             = 23
	CQLEXISTS                    = 24
	CQLEXIST                     = 25
	CQLDOES                      = 26
	CQLIN                        = 27
	CQLPOINT                     = 28
	CQLLINESTRING                = 29
	CQLPOLYGON                   = 30
	CQLMULTIPOINT                = 31
	CQLMULTILINESTRING           = 32
	CQLMULTIPOLYGON              = 33
	CQLGEOMETRYCOLLECTION        = 34
	CQLENVELOPE                  = 35
	CQLNumericLiteral            = 36
	CQLIdentifier                = 37
	CQLIdentifierStart           = 38
	CQLIdentifierPart            = 39
	CQLALPHA                     = 40
	CQLDIGIT                     = 41
	CQLOCTOTHORP                 = 42
	CQLDOLLAR                    = 43
	CQLUNDERSCORE                = 44
	CQLDOUBLEQUOTE               = 45
	CQLPERCENT                   = 46
	CQLAMPERSAND                 = 47
	CQLQUOTE                     = 48
	CQLLEFTPAREN                 = 49
	CQLRIGHTPAREN                = 50
	CQLLEFTSQUAREBRACKET         = 51
	CQLRIGHTSQUAREBRACKET        = 52
	CQLASTERISK                  = 53
	CQLPLUS                      = 54
	CQLCOMMA                     = 55
	CQLMINUS                     = 56
	CQLPERIOD                    = 57
	CQLSOLIDUS                   = 58
	CQLCOLON                     = 59
	CQLSEMICOLON                 = 60
	CQLQUESTIONMARK              = 61
	CQLVERTICALBAR               = 62
	CQLBIT                       = 63
	CQLHEXIT                     = 64
	CQLUnsignedNumericLiteral    = 65
	CQLSignedNumericLiteral      = 66
	CQLExactNumericLiteral       = 67
	CQLApproximateNumericLiteral = 68
	CQLMantissa                  = 69
	CQLExponent                  = 70
	CQLSignedInteger             = 71
	CQLUnsignedInteger           = 72
	CQLSign                      = 73
	CQLTemporalLiteral           = 74
	CQLInstant                   = 75
	CQLInterval                  = 76
	CQLInstantInInterval         = 77
	CQLFullDate                  = 78
	CQLDateYear                  = 79
	CQLDateMonth                 = 80
	CQLDateDay                   = 81
	CQLUtcTime                   = 82
	CQLTimeZoneOffset            = 83
	CQLTimeHour                  = 84
	CQLTimeMinute                = 85
	CQLTimeSecond                = 86
	CQLNOW                       = 87
	CQLWS                        = 88
	CQLCharacterStringLiteral    = 89
	CQLQuotedQuote               = 90
)

// CQL rules.
const (
	CQLRULE_cqlFilter                 = 0
	CQLRULE_booleanValueExpression    = 1
	CQLRULE_booleanTerm               = 2
	CQLRULE_booleanFactor             = 3
	CQLRULE_booleanPrimary            = 4
	CQLRULE_predicate                 = 5
	CQLRULE_binaryComparisonPredicate = 6
	CQLRULE_likePredicate             = 7
	CQLRULE_betweenPredicate          = 8
	CQLRULE_isNullPredicate           = 9
	CQLRULE_scalarExpression          = 10
	CQLRULE_propertyName              = 11
	CQLRULE_characterLiteral          = 12
	CQLRULE_numericLiteral            = 13
	CQLRULE_booleanLiteral            = 14
	CQLRULE_spatialPredicate          = 15
	CQLRULE_geomExpression            = 16
	CQLRULE_geomLiteral               = 17
	CQLRULE_point                     = 18
	CQLRULE_linestring                = 19
	CQLRULE_linestringDef             = 20
	CQLRULE_polygon                   = 21
	CQLRULE_polygonDef                = 22
	CQLRULE_multiPoint                = 23
	CQLRULE_multiLinestring           = 24
	CQLRULE_multiPolygon              = 25
	CQLRULE_geometryCollection        = 26
	CQLRULE_envelope                  = 27
	CQLRULE_coordinate                = 28
	CQLRULE_xCoord                    = 29
	CQLRULE_yCoord                    = 30
	CQLRULE_zCoord                    = 31
	CQLRULE_westBoundLon              = 32
	CQLRULE_eastBoundLon              = 33
	CQLRULE_northBoundLat             = 34
	CQLRULE_southBoundLat             = 35
	CQLRULE_minElev                   = 36
	CQLRULE_maxElev                   = 37
	CQLRULE_temporalPredicate         = 38
	CQLRULE_temporalExpression        = 39
	CQLRULE_temporalLiteral           = 40
	CQLRULE_inPredicate               = 41
)

// ICqlFilterContext is an interface to support dynamic dispatch.
type ICqlFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCqlFilterContext differentiates from other interfaces.
	IsCqlFilterContext()
}

type CqlFilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlFilterContext() *CqlFilterContext {
	var p = new(CqlFilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_cqlFilter
	return p
}

func (*CqlFilterContext) IsCqlFilterContext() {}

func NewCqlFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlFilterContext {
	var p = new(CqlFilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_cqlFilter

	return p
}

func (s *CqlFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlFilterContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *CqlFilterContext) EOF() antlr.TerminalNode {
	return s.GetToken(CQLEOF, 0)
}

func (s *CqlFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCqlFilter(s)
	}
}

func (s *CqlFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCqlFilter(s)
	}
}

func (p *CQL) CqlFilter() (localctx ICqlFilterContext) {
	localctx = NewCqlFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQLRULE_cqlFilter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.booleanValueExpression(0)
	}
	{
		p.SetState(85)
		p.Match(CQLEOF)
	}

	return localctx
}

// IBooleanValueExpressionContext is an interface to support dynamic dispatch.
type IBooleanValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueExpressionContext differentiates from other interfaces.
	IsBooleanValueExpressionContext()
}

type BooleanValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueExpressionContext() *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_booleanValueExpression
	return p
}

func (*BooleanValueExpressionContext) IsBooleanValueExpressionContext() {}

func NewBooleanValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueExpressionContext {
	var p = new(BooleanValueExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_booleanValueExpression

	return p
}

func (s *BooleanValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueExpressionContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanValueExpressionContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanValueExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(CQLOR, 0)
}

func (s *BooleanValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBooleanValueExpression(s)
	}
}

func (s *BooleanValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBooleanValueExpression(s)
	}
}

func (p *CQL) BooleanValueExpression() (localctx IBooleanValueExpressionContext) {
	return p.booleanValueExpression(0)
}

func (p *CQL) booleanValueExpression(_p int) (localctx IBooleanValueExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CQLRULE_booleanValueExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.booleanTerm(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanValueExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLRULE_booleanValueExpression)
			p.SetState(90)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(91)
				p.Match(CQLOR)
			}
			{
				p.SetState(92)
				p.booleanTerm(0)
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanTermContext is an interface to support dynamic dispatch.
type IBooleanTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanTermContext differentiates from other interfaces.
	IsBooleanTermContext()
}

type BooleanTermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanTermContext() *BooleanTermContext {
	var p = new(BooleanTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_booleanTerm
	return p
}

func (*BooleanTermContext) IsBooleanTermContext() {}

func NewBooleanTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTermContext {
	var p = new(BooleanTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_booleanTerm

	return p
}

func (s *BooleanTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTermContext) BooleanFactor() IBooleanFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanFactorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanFactorContext)
}

func (s *BooleanTermContext) BooleanTerm() IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanTermContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLAND, 0)
}

func (s *BooleanTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBooleanTerm(s)
	}
}

func (s *BooleanTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBooleanTerm(s)
	}
}

func (p *CQL) BooleanTerm() (localctx IBooleanTermContext) {
	return p.booleanTerm(0)
}

func (p *CQL) booleanTerm(_p int) (localctx IBooleanTermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanTermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CQLRULE_booleanTerm, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.BooleanFactor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBooleanTermContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLRULE_booleanTerm)
			p.SetState(101)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(102)
				p.Match(CQLAND)
			}
			{
				p.SetState(103)
				p.BooleanFactor()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanFactorContext is an interface to support dynamic dispatch.
type IBooleanFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanFactorContext differentiates from other interfaces.
	IsBooleanFactorContext()
}

type BooleanFactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanFactorContext() *BooleanFactorContext {
	var p = new(BooleanFactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_booleanFactor
	return p
}

func (*BooleanFactorContext) IsBooleanFactorContext() {}

func NewBooleanFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFactorContext {
	var p = new(BooleanFactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_booleanFactor

	return p
}

func (s *BooleanFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFactorContext) BooleanPrimary() IBooleanPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanPrimaryContext)
}

func (s *BooleanFactorContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLNOT, 0)
}

func (s *BooleanFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBooleanFactor(s)
	}
}

func (s *BooleanFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBooleanFactor(s)
	}
}

func (p *CQL) BooleanFactor() (localctx IBooleanFactorContext) {
	localctx = NewBooleanFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQLRULE_booleanFactor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(109)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(112)
		p.BooleanPrimary()
	}

	return localctx
}

// IBooleanPrimaryContext is an interface to support dynamic dispatch.
type IBooleanPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanPrimaryContext differentiates from other interfaces.
	IsBooleanPrimaryContext()
}

type BooleanPrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanPrimaryContext() *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_booleanPrimary
	return p
}

func (*BooleanPrimaryContext) IsBooleanPrimaryContext() {}

func NewBooleanPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_booleanPrimary

	return p
}

func (s *BooleanPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanPrimaryContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BooleanPrimaryContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *BooleanPrimaryContext) BooleanValueExpression() IBooleanValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueExpressionContext)
}

func (s *BooleanPrimaryContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *BooleanPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBooleanPrimary(s)
	}
}

func (s *BooleanPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBooleanPrimary(s)
	}
}

func (p *CQL) BooleanPrimary() (localctx IBooleanPrimaryContext) {
	localctx = NewBooleanPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQLRULE_booleanPrimary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLBooleanLiteral, CQLNumericLiteral, CQLIdentifier, CQLCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Predicate()
		}

	case CQLLEFTPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(CQLLEFTPAREN)
		}
		{
			p.SetState(116)
			p.booleanValueExpression(0)
		}
		{
			p.SetState(117)
			p.Match(CQLRIGHTPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) BinaryComparisonPredicate() IBinaryComparisonPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryComparisonPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryComparisonPredicateContext)
}

func (s *PredicateContext) LikePredicate() ILikePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikePredicateContext)
}

func (s *PredicateContext) BetweenPredicate() IBetweenPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBetweenPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBetweenPredicateContext)
}

func (s *PredicateContext) IsNullPredicate() IIsNullPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsNullPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsNullPredicateContext)
}

func (s *PredicateContext) InPredicate() IInPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInPredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *CQL) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQLRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.BinaryComparisonPredicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.LikePredicate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(123)
			p.BetweenPredicate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(124)
			p.IsNullPredicate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(125)
			p.InPredicate()
		}

	}

	return localctx
}

// IBinaryComparisonPredicateContext is an interface to support dynamic dispatch.
type IBinaryComparisonPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryComparisonPredicateContext differentiates from other interfaces.
	IsBinaryComparisonPredicateContext()
}

type BinaryComparisonPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryComparisonPredicateContext() *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_binaryComparisonPredicate
	return p
}

func (*BinaryComparisonPredicateContext) IsBinaryComparisonPredicateContext() {}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_binaryComparisonPredicate

	return p
}

func (s *BinaryComparisonPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryComparisonPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *BinaryComparisonPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BinaryComparisonPredicateContext) ComparisonOperator() antlr.TerminalNode {
	return s.GetToken(CQLComparisonOperator, 0)
}

func (s *BinaryComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (p *CQL) BinaryComparisonPredicate() (localctx IBinaryComparisonPredicateContext) {
	localctx = NewBinaryComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQLRULE_binaryComparisonPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.ScalarExpression()
	}
	{
		p.SetState(129)
		p.Match(CQLComparisonOperator)
	}
	{
		p.SetState(130)
		p.ScalarExpression()
	}

	return localctx
}

// ILikePredicateContext is an interface to support dynamic dispatch.
type ILikePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikePredicateContext differentiates from other interfaces.
	IsLikePredicateContext()
}

type LikePredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikePredicateContext() *LikePredicateContext {
	var p = new(LikePredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_likePredicate
	return p
}

func (*LikePredicateContext) IsLikePredicateContext() {}

func NewLikePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_likePredicate

	return p
}

func (s *LikePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *LikePredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *LikePredicateContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(CQLLIKE, 0)
}

func (s *LikePredicateContext) ILIKE() antlr.TerminalNode {
	return s.GetToken(CQLILIKE, 0)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLNOT, 0)
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (p *CQL) LikePredicate() (localctx ILikePredicateContext) {
	localctx = NewLikePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLRULE_likePredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.PropertyName()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(133)
			p.Match(CQLNOT)
		}

	}
	p.SetState(136)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLLIKE || _la == CQLILIKE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(137)
		p.CharacterLiteral()
	}

	return localctx
}

// IBetweenPredicateContext is an interface to support dynamic dispatch.
type IBetweenPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBetweenPredicateContext differentiates from other interfaces.
	IsBetweenPredicateContext()
}

type BetweenPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBetweenPredicateContext() *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_betweenPredicate
	return p
}

func (*BetweenPredicateContext) IsBetweenPredicateContext() {}

func NewBetweenPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_betweenPredicate

	return p
}

func (s *BetweenPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(CQLBETWEEN, 0)
}

func (s *BetweenPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *BetweenPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLAND, 0)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLNOT, 0)
}

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBetweenPredicate(s)
	}
}

func (p *CQL) BetweenPredicate() (localctx IBetweenPredicateContext) {
	localctx = NewBetweenPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQLRULE_betweenPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.PropertyName()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(140)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(143)
		p.Match(CQLBETWEEN)
	}
	{
		p.SetState(144)
		p.ScalarExpression()
	}
	{
		p.SetState(145)
		p.Match(CQLAND)
	}
	{
		p.SetState(146)
		p.ScalarExpression()
	}

	return localctx
}

// IIsNullPredicateContext is an interface to support dynamic dispatch.
type IIsNullPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsNullPredicateContext differentiates from other interfaces.
	IsIsNullPredicateContext()
}

type IsNullPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsNullPredicateContext() *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_isNullPredicate
	return p
}

func (*IsNullPredicateContext) IsIsNullPredicateContext() {}

func NewIsNullPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_isNullPredicate

	return p
}

func (s *IsNullPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsNullPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(CQLIS, 0)
}

func (s *IsNullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(CQLNULL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLNOT, 0)
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (p *CQL) IsNullPredicate() (localctx IIsNullPredicateContext) {
	localctx = NewIsNullPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLRULE_isNullPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.PropertyName()
	}
	{
		p.SetState(149)
		p.Match(CQLIS)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(150)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(153)
		p.Match(CQLNULL)
	}

	return localctx
}

// IScalarExpressionContext is an interface to support dynamic dispatch.
type IScalarExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarExpressionContext differentiates from other interfaces.
	IsScalarExpressionContext()
}

type ScalarExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarExpressionContext() *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_scalarExpression
	return p
}

func (*ScalarExpressionContext) IsScalarExpressionContext() {}

func NewScalarExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_scalarExpression

	return p
}

func (s *ScalarExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarExpressionContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *ScalarExpressionContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *ScalarExpressionContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *ScalarExpressionContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ScalarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterScalarExpression(s)
	}
}

func (s *ScalarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitScalarExpression(s)
	}
}

func (p *CQL) ScalarExpression() (localctx IScalarExpressionContext) {
	localctx = NewScalarExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CQLRULE_scalarExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.PropertyName()
		}

	case CQLCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.CharacterLiteral()
		}

	case CQLNumericLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.NumericLiteral()
		}

	case CQLBooleanLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.BooleanLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLIdentifier, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *CQL) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CQLRULE_propertyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(CQLIdentifier)
	}

	return localctx
}

// ICharacterLiteralContext is an interface to support dynamic dispatch.
type ICharacterLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharacterLiteralContext differentiates from other interfaces.
	IsCharacterLiteralContext()
}

type CharacterLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharacterLiteralContext() *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_characterLiteral
	return p
}

func (*CharacterLiteralContext) IsCharacterLiteralContext() {}

func NewCharacterLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_characterLiteral

	return p
}

func (s *CharacterLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterLiteralContext) CharacterStringLiteral() antlr.TerminalNode {
	return s.GetToken(CQLCharacterStringLiteral, 0)
}

func (s *CharacterLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCharacterLiteral(s)
	}
}

func (s *CharacterLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCharacterLiteral(s)
	}
}

func (p *CQL) CharacterLiteral() (localctx ICharacterLiteralContext) {
	localctx = NewCharacterLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CQLRULE_characterLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(CQLCharacterStringLiteral)
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *CQL) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQLRULE_numericLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(CQLBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *CQL) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQLRULE_booleanLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(CQLBooleanLiteral)
	}

	return localctx
}

// ISpatialPredicateContext is an interface to support dynamic dispatch.
type ISpatialPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpatialPredicateContext differentiates from other interfaces.
	IsSpatialPredicateContext()
}

type SpatialPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpatialPredicateContext() *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_spatialPredicate
	return p
}

func (*SpatialPredicateContext) IsSpatialPredicateContext() {}

func NewSpatialPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_spatialPredicate

	return p
}

func (s *SpatialPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpatialPredicateContext) SpatialOperator() antlr.TerminalNode {
	return s.GetToken(CQLSpatialOperator, 0)
}

func (s *SpatialPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *SpatialPredicateContext) AllGeomExpression() []IGeomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem())
	var tst = make([]IGeomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomExpressionContext)
		}
	}

	return tst
}

func (s *SpatialPredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *SpatialPredicateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, 0)
}

func (s *SpatialPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *SpatialPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpatialPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpatialPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterSpatialPredicate(s)
	}
}

func (s *SpatialPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitSpatialPredicate(s)
	}
}

func (p *CQL) SpatialPredicate() (localctx ISpatialPredicateContext) {
	localctx = NewSpatialPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQLRULE_spatialPredicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(CQLSpatialOperator)
	}
	{
		p.SetState(170)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(171)
		p.GeomExpression()
	}
	{
		p.SetState(172)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(173)
		p.GeomExpression()
	}
	{
		p.SetState(174)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IGeomExpressionContext is an interface to support dynamic dispatch.
type IGeomExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeomExpressionContext differentiates from other interfaces.
	IsGeomExpressionContext()
}

type GeomExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeomExpressionContext() *GeomExpressionContext {
	var p = new(GeomExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_geomExpression
	return p
}

func (*GeomExpressionContext) IsGeomExpressionContext() {}

func NewGeomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomExpressionContext {
	var p = new(GeomExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_geomExpression

	return p
}

func (s *GeomExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomExpressionContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GeomExpressionContext) GeomLiteral() IGeomLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterGeomExpression(s)
	}
}

func (s *GeomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitGeomExpression(s)
	}
}

func (p *CQL) GeomExpression() (localctx IGeomExpressionContext) {
	localctx = NewGeomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLRULE_geomExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.PropertyName()
		}

	case CQLPOINT, CQLLINESTRING, CQLPOLYGON, CQLMULTIPOINT, CQLMULTILINESTRING, CQLMULTIPOLYGON, CQLGEOMETRYCOLLECTION, CQLENVELOPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.GeomLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGeomLiteralContext is an interface to support dynamic dispatch.
type IGeomLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeomLiteralContext differentiates from other interfaces.
	IsGeomLiteralContext()
}

type GeomLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeomLiteralContext() *GeomLiteralContext {
	var p = new(GeomLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_geomLiteral
	return p
}

func (*GeomLiteralContext) IsGeomLiteralContext() {}

func NewGeomLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomLiteralContext {
	var p = new(GeomLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_geomLiteral

	return p
}

func (s *GeomLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *GeomLiteralContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *GeomLiteralContext) Linestring() ILinestringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinestringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinestringContext)
}

func (s *GeomLiteralContext) Polygon() IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *GeomLiteralContext) MultiPoint() IMultiPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPointContext)
}

func (s *GeomLiteralContext) MultiLinestring() IMultiLinestringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiLinestringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiLinestringContext)
}

func (s *GeomLiteralContext) MultiPolygon() IMultiPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPolygonContext)
}

func (s *GeomLiteralContext) GeometryCollection() IGeometryCollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeometryCollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeometryCollectionContext)
}

func (s *GeomLiteralContext) Envelope() IEnvelopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnvelopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnvelopeContext)
}

func (s *GeomLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeomLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeomLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterGeomLiteral(s)
	}
}

func (s *GeomLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitGeomLiteral(s)
	}
}

func (p *CQL) GeomLiteral() (localctx IGeomLiteralContext) {
	localctx = NewGeomLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQLRULE_geomLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLPOINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Point()
		}

	case CQLLINESTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Linestring()
		}

	case CQLPOLYGON:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Polygon()
		}

	case CQLMULTIPOINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.MultiPoint()
		}

	case CQLMULTILINESTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.MultiLinestring()
		}

	case CQLMULTIPOLYGON:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)
			p.MultiPolygon()
		}

	case CQLGEOMETRYCOLLECTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(186)
			p.GeometryCollection()
		}

	case CQLENVELOPE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(187)
			p.Envelope()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_point
	return p
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) POINT() antlr.TerminalNode {
	return s.GetToken(CQLPOINT, 0)
}

func (s *PointContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *PointContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *PointContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *CQL) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQLRULE_point)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(CQLPOINT)
	}
	{
		p.SetState(191)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(192)
		p.Coordinate()
	}
	{
		p.SetState(193)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// ILinestringContext is an interface to support dynamic dispatch.
type ILinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinestringContext differentiates from other interfaces.
	IsLinestringContext()
}

type LinestringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinestringContext() *LinestringContext {
	var p = new(LinestringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_linestring
	return p
}

func (*LinestringContext) IsLinestringContext() {}

func NewLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinestringContext {
	var p = new(LinestringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_linestring

	return p
}

func (s *LinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *LinestringContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLLINESTRING, 0)
}

func (s *LinestringContext) LinestringDef() ILinestringDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinestringDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILinestringDefContext)
}

func (s *LinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterLinestring(s)
	}
}

func (s *LinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitLinestring(s)
	}
}

func (p *CQL) Linestring() (localctx ILinestringContext) {
	localctx = NewLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CQLRULE_linestring)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(CQLLINESTRING)
	}
	{
		p.SetState(196)
		p.LinestringDef()
	}

	return localctx
}

// ILinestringDefContext is an interface to support dynamic dispatch.
type ILinestringDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinestringDefContext differentiates from other interfaces.
	IsLinestringDefContext()
}

type LinestringDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinestringDefContext() *LinestringDefContext {
	var p = new(LinestringDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_linestringDef
	return p
}

func (*LinestringDefContext) IsLinestringDefContext() {}

func NewLinestringDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinestringDefContext {
	var p = new(LinestringDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_linestringDef

	return p
}

func (s *LinestringDefContext) GetParser() antlr.Parser { return s.parser }

func (s *LinestringDefContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *LinestringDefContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *LinestringDefContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *LinestringDefContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *LinestringDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *LinestringDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *LinestringDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinestringDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinestringDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterLinestringDef(s)
	}
}

func (s *LinestringDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitLinestringDef(s)
	}
}

func (p *CQL) LinestringDef() (localctx ILinestringDefContext) {
	localctx = NewLinestringDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CQLRULE_linestringDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(199)
		p.Coordinate()
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(200)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(201)
			p.Coordinate()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(207)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IPolygonContext is an interface to support dynamic dispatch.
type IPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonContext differentiates from other interfaces.
	IsPolygonContext()
}

type PolygonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_polygon
	return p
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(CQLPOLYGON, 0)
}

func (s *PolygonContext) PolygonDef() IPolygonDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *PolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *CQL) Polygon() (localctx IPolygonContext) {
	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQLRULE_polygon)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(CQLPOLYGON)
	}
	{
		p.SetState(210)
		p.PolygonDef()
	}

	return localctx
}

// IPolygonDefContext is an interface to support dynamic dispatch.
type IPolygonDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonDefContext differentiates from other interfaces.
	IsPolygonDefContext()
}

type PolygonDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonDefContext() *PolygonDefContext {
	var p = new(PolygonDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_polygonDef
	return p
}

func (*PolygonDefContext) IsPolygonDefContext() {}

func NewPolygonDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonDefContext {
	var p = new(PolygonDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_polygonDef

	return p
}

func (s *PolygonDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonDefContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *PolygonDefContext) AllLinestringDef() []ILinestringDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILinestringDefContext)(nil)).Elem())
	var tst = make([]ILinestringDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILinestringDefContext)
		}
	}

	return tst
}

func (s *PolygonDefContext) LinestringDef(i int) ILinestringDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinestringDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILinestringDefContext)
}

func (s *PolygonDefContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *PolygonDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *PolygonDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *PolygonDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterPolygonDef(s)
	}
}

func (s *PolygonDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitPolygonDef(s)
	}
}

func (p *CQL) PolygonDef() (localctx IPolygonDefContext) {
	localctx = NewPolygonDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CQLRULE_polygonDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(213)
		p.LinestringDef()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(214)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(215)
			p.LinestringDef()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IMultiPointContext is an interface to support dynamic dispatch.
type IMultiPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPointContext differentiates from other interfaces.
	IsMultiPointContext()
}

type MultiPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPointContext() *MultiPointContext {
	var p = new(MultiPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_multiPoint
	return p
}

func (*MultiPointContext) IsMultiPointContext() {}

func NewMultiPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointContext {
	var p = new(MultiPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_multiPoint

	return p
}

func (s *MultiPointContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(CQLMULTIPOINT, 0)
}

func (s *MultiPointContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *MultiPointContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *MultiPointContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *MultiPointContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *MultiPointContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *MultiPointContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *MultiPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterMultiPoint(s)
	}
}

func (s *MultiPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitMultiPoint(s)
	}
}

func (p *CQL) MultiPoint() (localctx IMultiPointContext) {
	localctx = NewMultiPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CQLRULE_multiPoint)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(CQLMULTIPOINT)
	}
	{
		p.SetState(224)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(225)
		p.Coordinate()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(226)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(227)
			p.Coordinate()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IMultiLinestringContext is an interface to support dynamic dispatch.
type IMultiLinestringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiLinestringContext differentiates from other interfaces.
	IsMultiLinestringContext()
}

type MultiLinestringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiLinestringContext() *MultiLinestringContext {
	var p = new(MultiLinestringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_multiLinestring
	return p
}

func (*MultiLinestringContext) IsMultiLinestringContext() {}

func NewMultiLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLinestringContext {
	var p = new(MultiLinestringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_multiLinestring

	return p
}

func (s *MultiLinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLinestringContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLMULTILINESTRING, 0)
}

func (s *MultiLinestringContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *MultiLinestringContext) AllLinestringDef() []ILinestringDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILinestringDefContext)(nil)).Elem())
	var tst = make([]ILinestringDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILinestringDefContext)
		}
	}

	return tst
}

func (s *MultiLinestringContext) LinestringDef(i int) ILinestringDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinestringDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILinestringDefContext)
}

func (s *MultiLinestringContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *MultiLinestringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *MultiLinestringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *MultiLinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterMultiLinestring(s)
	}
}

func (s *MultiLinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitMultiLinestring(s)
	}
}

func (p *CQL) MultiLinestring() (localctx IMultiLinestringContext) {
	localctx = NewMultiLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CQLRULE_multiLinestring)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(CQLMULTILINESTRING)
	}
	{
		p.SetState(236)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(237)
		p.LinestringDef()
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(238)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(239)
			p.LinestringDef()
		}

		p.SetState(244)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(245)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IMultiPolygonContext is an interface to support dynamic dispatch.
type IMultiPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPolygonContext differentiates from other interfaces.
	IsMultiPolygonContext()
}

type MultiPolygonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonContext() *MultiPolygonContext {
	var p = new(MultiPolygonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_multiPolygon
	return p
}

func (*MultiPolygonContext) IsMultiPolygonContext() {}

func NewMultiPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonContext {
	var p = new(MultiPolygonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_multiPolygon

	return p
}

func (s *MultiPolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(CQLMULTIPOLYGON, 0)
}

func (s *MultiPolygonContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *MultiPolygonContext) AllPolygonDef() []IPolygonDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem())
	var tst = make([]IPolygonDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonDefContext)
		}
	}

	return tst
}

func (s *MultiPolygonContext) PolygonDef(i int) IPolygonDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonDefContext)
}

func (s *MultiPolygonContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *MultiPolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *MultiPolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *MultiPolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterMultiPolygon(s)
	}
}

func (s *MultiPolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitMultiPolygon(s)
	}
}

func (p *CQL) MultiPolygon() (localctx IMultiPolygonContext) {
	localctx = NewMultiPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CQLRULE_multiPolygon)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(CQLMULTIPOLYGON)
	}
	{
		p.SetState(248)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(249)
		p.PolygonDef()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(250)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(251)
			p.PolygonDef()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IGeometryCollectionContext is an interface to support dynamic dispatch.
type IGeometryCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryCollectionContext differentiates from other interfaces.
	IsGeometryCollectionContext()
}

type GeometryCollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeometryCollectionContext() *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_geometryCollection
	return p
}

func (*GeometryCollectionContext) IsGeometryCollectionContext() {}

func NewGeometryCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_geometryCollection

	return p
}

func (s *GeometryCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryCollectionContext) GEOMETRYCOLLECTION() antlr.TerminalNode {
	return s.GetToken(CQLGEOMETRYCOLLECTION, 0)
}

func (s *GeometryCollectionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *GeometryCollectionContext) AllGeomLiteral() []IGeomLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem())
	var tst = make([]IGeomLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomLiteralContext)
		}
	}

	return tst
}

func (s *GeometryCollectionContext) GeomLiteral(i int) IGeomLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomLiteralContext)
}

func (s *GeometryCollectionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *GeometryCollectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *GeometryCollectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *GeometryCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterGeometryCollection(s)
	}
}

func (s *GeometryCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitGeometryCollection(s)
	}
}

func (p *CQL) GeometryCollection() (localctx IGeometryCollectionContext) {
	localctx = NewGeometryCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CQLRULE_geometryCollection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(CQLGEOMETRYCOLLECTION)
	}
	{
		p.SetState(260)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(261)
		p.GeomLiteral()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(262)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(263)
			p.GeomLiteral()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(269)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// IEnvelopeContext is an interface to support dynamic dispatch.
type IEnvelopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnvelopeContext differentiates from other interfaces.
	IsEnvelopeContext()
}

type EnvelopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvelopeContext() *EnvelopeContext {
	var p = new(EnvelopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_envelope
	return p
}

func (*EnvelopeContext) IsEnvelopeContext() {}

func NewEnvelopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvelopeContext {
	var p = new(EnvelopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_envelope

	return p
}

func (s *EnvelopeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvelopeContext) ENVELOPE() antlr.TerminalNode {
	return s.GetToken(CQLENVELOPE, 0)
}

func (s *EnvelopeContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *EnvelopeContext) WestBoundLon() IWestBoundLonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWestBoundLonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWestBoundLonContext)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *EnvelopeContext) SouthBoundLat() ISouthBoundLatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISouthBoundLatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISouthBoundLatContext)
}

func (s *EnvelopeContext) EastBoundLon() IEastBoundLonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEastBoundLonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEastBoundLonContext)
}

func (s *EnvelopeContext) NorthBoundLat() INorthBoundLatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INorthBoundLatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INorthBoundLatContext)
}

func (s *EnvelopeContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *EnvelopeContext) MinElev() IMinElevContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinElevContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinElevContext)
}

func (s *EnvelopeContext) MaxElev() IMaxElevContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMaxElevContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMaxElevContext)
}

func (s *EnvelopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvelopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvelopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterEnvelope(s)
	}
}

func (s *EnvelopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitEnvelope(s)
	}
}

func (p *CQL) Envelope() (localctx IEnvelopeContext) {
	localctx = NewEnvelopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CQLRULE_envelope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(CQLENVELOPE)
	}
	{
		p.SetState(272)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(273)
		p.WestBoundLon()
	}
	{
		p.SetState(274)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(275)
		p.SouthBoundLat()
	}
	{
		p.SetState(276)
		p.Match(CQLCOMMA)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(277)
			p.MinElev()
		}
		{
			p.SetState(278)
			p.Match(CQLCOMMA)
		}

	}
	{
		p.SetState(282)
		p.EastBoundLon()
	}
	{
		p.SetState(283)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(284)
		p.NorthBoundLat()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLCOMMA {
		{
			p.SetState(285)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(286)
			p.MaxElev()
		}

	}
	{
		p.SetState(289)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) XCoord() IXCoordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXCoordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IXCoordContext)
}

func (s *CoordinateContext) YCoord() IYCoordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYCoordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYCoordContext)
}

func (s *CoordinateContext) ZCoord() IZCoordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IZCoordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IZCoordContext)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCoordinate(s)
	}
}

func (s *CoordinateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCoordinate(s)
	}
}

func (p *CQL) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CQLRULE_coordinate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.XCoord()
	}
	{
		p.SetState(292)
		p.YCoord()
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNumericLiteral {
		{
			p.SetState(293)
			p.ZCoord()
		}

	}

	return localctx
}

// IXCoordContext is an interface to support dynamic dispatch.
type IXCoordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXCoordContext differentiates from other interfaces.
	IsXCoordContext()
}

type XCoordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXCoordContext() *XCoordContext {
	var p = new(XCoordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_xCoord
	return p
}

func (*XCoordContext) IsXCoordContext() {}

func NewXCoordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XCoordContext {
	var p = new(XCoordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_xCoord

	return p
}

func (s *XCoordContext) GetParser() antlr.Parser { return s.parser }

func (s *XCoordContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *XCoordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XCoordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XCoordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterXCoord(s)
	}
}

func (s *XCoordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitXCoord(s)
	}
}

func (p *CQL) XCoord() (localctx IXCoordContext) {
	localctx = NewXCoordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CQLRULE_xCoord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IYCoordContext is an interface to support dynamic dispatch.
type IYCoordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYCoordContext differentiates from other interfaces.
	IsYCoordContext()
}

type YCoordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYCoordContext() *YCoordContext {
	var p = new(YCoordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_yCoord
	return p
}

func (*YCoordContext) IsYCoordContext() {}

func NewYCoordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YCoordContext {
	var p = new(YCoordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_yCoord

	return p
}

func (s *YCoordContext) GetParser() antlr.Parser { return s.parser }

func (s *YCoordContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *YCoordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YCoordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YCoordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterYCoord(s)
	}
}

func (s *YCoordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitYCoord(s)
	}
}

func (p *CQL) YCoord() (localctx IYCoordContext) {
	localctx = NewYCoordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CQLRULE_yCoord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IZCoordContext is an interface to support dynamic dispatch.
type IZCoordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZCoordContext differentiates from other interfaces.
	IsZCoordContext()
}

type ZCoordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZCoordContext() *ZCoordContext {
	var p = new(ZCoordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_zCoord
	return p
}

func (*ZCoordContext) IsZCoordContext() {}

func NewZCoordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZCoordContext {
	var p = new(ZCoordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_zCoord

	return p
}

func (s *ZCoordContext) GetParser() antlr.Parser { return s.parser }

func (s *ZCoordContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *ZCoordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZCoordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZCoordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterZCoord(s)
	}
}

func (s *ZCoordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitZCoord(s)
	}
}

func (p *CQL) ZCoord() (localctx IZCoordContext) {
	localctx = NewZCoordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CQLRULE_zCoord)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IWestBoundLonContext is an interface to support dynamic dispatch.
type IWestBoundLonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWestBoundLonContext differentiates from other interfaces.
	IsWestBoundLonContext()
}

type WestBoundLonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWestBoundLonContext() *WestBoundLonContext {
	var p = new(WestBoundLonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_westBoundLon
	return p
}

func (*WestBoundLonContext) IsWestBoundLonContext() {}

func NewWestBoundLonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WestBoundLonContext {
	var p = new(WestBoundLonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_westBoundLon

	return p
}

func (s *WestBoundLonContext) GetParser() antlr.Parser { return s.parser }

func (s *WestBoundLonContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *WestBoundLonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WestBoundLonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WestBoundLonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterWestBoundLon(s)
	}
}

func (s *WestBoundLonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitWestBoundLon(s)
	}
}

func (p *CQL) WestBoundLon() (localctx IWestBoundLonContext) {
	localctx = NewWestBoundLonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CQLRULE_westBoundLon)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IEastBoundLonContext is an interface to support dynamic dispatch.
type IEastBoundLonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEastBoundLonContext differentiates from other interfaces.
	IsEastBoundLonContext()
}

type EastBoundLonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEastBoundLonContext() *EastBoundLonContext {
	var p = new(EastBoundLonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_eastBoundLon
	return p
}

func (*EastBoundLonContext) IsEastBoundLonContext() {}

func NewEastBoundLonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EastBoundLonContext {
	var p = new(EastBoundLonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_eastBoundLon

	return p
}

func (s *EastBoundLonContext) GetParser() antlr.Parser { return s.parser }

func (s *EastBoundLonContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *EastBoundLonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EastBoundLonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EastBoundLonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterEastBoundLon(s)
	}
}

func (s *EastBoundLonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitEastBoundLon(s)
	}
}

func (p *CQL) EastBoundLon() (localctx IEastBoundLonContext) {
	localctx = NewEastBoundLonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CQLRULE_eastBoundLon)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// INorthBoundLatContext is an interface to support dynamic dispatch.
type INorthBoundLatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNorthBoundLatContext differentiates from other interfaces.
	IsNorthBoundLatContext()
}

type NorthBoundLatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNorthBoundLatContext() *NorthBoundLatContext {
	var p = new(NorthBoundLatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_northBoundLat
	return p
}

func (*NorthBoundLatContext) IsNorthBoundLatContext() {}

func NewNorthBoundLatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NorthBoundLatContext {
	var p = new(NorthBoundLatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_northBoundLat

	return p
}

func (s *NorthBoundLatContext) GetParser() antlr.Parser { return s.parser }

func (s *NorthBoundLatContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *NorthBoundLatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NorthBoundLatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NorthBoundLatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterNorthBoundLat(s)
	}
}

func (s *NorthBoundLatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitNorthBoundLat(s)
	}
}

func (p *CQL) NorthBoundLat() (localctx INorthBoundLatContext) {
	localctx = NewNorthBoundLatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CQLRULE_northBoundLat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// ISouthBoundLatContext is an interface to support dynamic dispatch.
type ISouthBoundLatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSouthBoundLatContext differentiates from other interfaces.
	IsSouthBoundLatContext()
}

type SouthBoundLatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySouthBoundLatContext() *SouthBoundLatContext {
	var p = new(SouthBoundLatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_southBoundLat
	return p
}

func (*SouthBoundLatContext) IsSouthBoundLatContext() {}

func NewSouthBoundLatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SouthBoundLatContext {
	var p = new(SouthBoundLatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_southBoundLat

	return p
}

func (s *SouthBoundLatContext) GetParser() antlr.Parser { return s.parser }

func (s *SouthBoundLatContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *SouthBoundLatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SouthBoundLatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SouthBoundLatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterSouthBoundLat(s)
	}
}

func (s *SouthBoundLatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitSouthBoundLat(s)
	}
}

func (p *CQL) SouthBoundLat() (localctx ISouthBoundLatContext) {
	localctx = NewSouthBoundLatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CQLRULE_southBoundLat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IMinElevContext is an interface to support dynamic dispatch.
type IMinElevContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinElevContext differentiates from other interfaces.
	IsMinElevContext()
}

type MinElevContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinElevContext() *MinElevContext {
	var p = new(MinElevContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_minElev
	return p
}

func (*MinElevContext) IsMinElevContext() {}

func NewMinElevContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinElevContext {
	var p = new(MinElevContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_minElev

	return p
}

func (s *MinElevContext) GetParser() antlr.Parser { return s.parser }

func (s *MinElevContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *MinElevContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinElevContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinElevContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterMinElev(s)
	}
}

func (s *MinElevContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitMinElev(s)
	}
}

func (p *CQL) MinElev() (localctx IMinElevContext) {
	localctx = NewMinElevContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CQLRULE_minElev)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// IMaxElevContext is an interface to support dynamic dispatch.
type IMaxElevContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaxElevContext differentiates from other interfaces.
	IsMaxElevContext()
}

type MaxElevContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxElevContext() *MaxElevContext {
	var p = new(MaxElevContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_maxElev
	return p
}

func (*MaxElevContext) IsMaxElevContext() {}

func NewMaxElevContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxElevContext {
	var p = new(MaxElevContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_maxElev

	return p
}

func (s *MaxElevContext) GetParser() antlr.Parser { return s.parser }

func (s *MaxElevContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *MaxElevContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxElevContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaxElevContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterMaxElev(s)
	}
}

func (s *MaxElevContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitMaxElev(s)
	}
}

func (p *CQL) MaxElev() (localctx IMaxElevContext) {
	localctx = NewMaxElevContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CQLRULE_maxElev)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(CQLNumericLiteral)
	}

	return localctx
}

// ITemporalPredicateContext is an interface to support dynamic dispatch.
type ITemporalPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemporalPredicateContext differentiates from other interfaces.
	IsTemporalPredicateContext()
}

type TemporalPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemporalPredicateContext() *TemporalPredicateContext {
	var p = new(TemporalPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_temporalPredicate
	return p
}

func (*TemporalPredicateContext) IsTemporalPredicateContext() {}

func NewTemporalPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalPredicateContext {
	var p = new(TemporalPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_temporalPredicate

	return p
}

func (s *TemporalPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalPredicateContext) AllTemporalExpression() []ITemporalExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITemporalExpressionContext)(nil)).Elem())
	var tst = make([]ITemporalExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITemporalExpressionContext)
		}
	}

	return tst
}

func (s *TemporalPredicateContext) TemporalExpression(i int) ITemporalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITemporalExpressionContext)
}

func (s *TemporalPredicateContext) TemporalOperator() antlr.TerminalNode {
	return s.GetToken(CQLTemporalOperator, 0)
}

func (s *TemporalPredicateContext) ComparisonOperator() antlr.TerminalNode {
	return s.GetToken(CQLComparisonOperator, 0)
}

func (s *TemporalPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterTemporalPredicate(s)
	}
}

func (s *TemporalPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitTemporalPredicate(s)
	}
}

func (p *CQL) TemporalPredicate() (localctx ITemporalPredicateContext) {
	localctx = NewTemporalPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CQLRULE_temporalPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.TemporalExpression()
	}
	p.SetState(315)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLComparisonOperator || _la == CQLTemporalOperator) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(316)
		p.TemporalExpression()
	}

	return localctx
}

// ITemporalExpressionContext is an interface to support dynamic dispatch.
type ITemporalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemporalExpressionContext differentiates from other interfaces.
	IsTemporalExpressionContext()
}

type TemporalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemporalExpressionContext() *TemporalExpressionContext {
	var p = new(TemporalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_temporalExpression
	return p
}

func (*TemporalExpressionContext) IsTemporalExpressionContext() {}

func NewTemporalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalExpressionContext {
	var p = new(TemporalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_temporalExpression

	return p
}

func (s *TemporalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalExpressionContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *TemporalExpressionContext) TemporalLiteral() ITemporalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporalLiteralContext)
}

func (s *TemporalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterTemporalExpression(s)
	}
}

func (s *TemporalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitTemporalExpression(s)
	}
}

func (p *CQL) TemporalExpression() (localctx ITemporalExpressionContext) {
	localctx = NewTemporalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CQLRULE_temporalExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.PropertyName()
		}

	case CQLTemporalLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.TemporalLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITemporalLiteralContext is an interface to support dynamic dispatch.
type ITemporalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemporalLiteralContext differentiates from other interfaces.
	IsTemporalLiteralContext()
}

type TemporalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemporalLiteralContext() *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_temporalLiteral
	return p
}

func (*TemporalLiteralContext) IsTemporalLiteralContext() {}

func NewTemporalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_temporalLiteral

	return p
}

func (s *TemporalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalLiteralContext) TemporalLiteral() antlr.TerminalNode {
	return s.GetToken(CQLTemporalLiteral, 0)
}

func (s *TemporalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterTemporalLiteral(s)
	}
}

func (s *TemporalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitTemporalLiteral(s)
	}
}

func (p *CQL) TemporalLiteral() (localctx ITemporalLiteralContext) {
	localctx = NewTemporalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CQLRULE_temporalLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(CQLTemporalLiteral)
	}

	return localctx
}

// IInPredicateContext is an interface to support dynamic dispatch.
type IInPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInPredicateContext differentiates from other interfaces.
	IsInPredicateContext()
}

type InPredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInPredicateContext() *InPredicateContext {
	var p = new(InPredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_inPredicate
	return p
}

func (*InPredicateContext) IsInPredicateContext() {}

func NewInPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InPredicateContext {
	var p = new(InPredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_inPredicate

	return p
}

func (s *InPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *InPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(CQLIN, 0)
}

func (s *InPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
}

func (s *InPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *InPredicateContext) AllCharacterLiteral() []ICharacterLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem())
	var tst = make([]ICharacterLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterLiteralContext)
		}
	}

	return tst
}

func (s *InPredicateContext) CharacterLiteral(i int) ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *InPredicateContext) AllNumericLiteral() []INumericLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem())
	var tst = make([]INumericLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralContext)
		}
	}

	return tst
}

func (s *InPredicateContext) NumericLiteral(i int) INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLNOT, 0)
}

func (s *InPredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *InPredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterInPredicate(s)
	}
}

func (s *InPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitInPredicate(s)
	}
}

func (p *CQL) InPredicate() (localctx IInPredicateContext) {
	localctx = NewInPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CQLRULE_inPredicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.PropertyName()
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(325)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(328)
		p.Match(CQLIN)
	}
	{
		p.SetState(329)
		p.Match(CQLLEFTPAREN)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLCharacterStringLiteral:
		{
			p.SetState(330)
			p.CharacterLiteral()
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLCOMMA {
			{
				p.SetState(331)
				p.Match(CQLCOMMA)
			}
			{
				p.SetState(332)
				p.CharacterLiteral()
			}

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case CQLNumericLiteral:
		{
			p.SetState(338)
			p.NumericLiteral()
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLCOMMA {
			{
				p.SetState(339)
				p.Match(CQLCOMMA)
			}
			{
				p.SetState(340)
				p.NumericLiteral()
			}

			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(348)
		p.Match(CQLRIGHTPAREN)
	}

	return localctx
}

func (p *CQL) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *BooleanValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanValueExpressionContext)
		}
		return p.BooleanValueExpression_Sempred(t, predIndex)

	case 2:
		var t *BooleanTermContext = nil
		if localctx != nil {
			t = localctx.(*BooleanTermContext)
		}
		return p.BooleanTerm_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CQL) BooleanValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CQL) BooleanTerm_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
