// Generated from CQLParser.g4 by ANTLR 4.7.

package cql // CQLParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 85, 324,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	5, 3, 79, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 84, 10, 4, 3, 5, 5, 5, 87, 10,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 97, 10, 6, 3,
	7, 3, 7, 3, 7, 5, 7, 102, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 109,
	10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 117, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 5, 11, 124, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 5, 12, 133, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 140, 10, 12, 12, 12, 14, 12, 143, 11, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 148, 10, 12, 12, 12, 14, 12, 151, 11, 12, 5, 12, 153, 10, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 160, 10, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 170, 10, 14, 3, 14, 3, 14,
	3, 14, 7, 14, 175, 10, 14, 12, 14, 14, 14, 178, 11, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 185, 10, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 5, 23, 215, 10, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 225, 10, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	29, 3, 29, 7, 29, 244, 10, 29, 12, 29, 14, 29, 247, 11, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 256, 10, 30, 12, 30, 14, 30,
	259, 11, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 268,
	10, 31, 12, 31, 14, 31, 271, 11, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 7, 32, 280, 10, 32, 12, 32, 14, 32, 283, 11, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 292, 10, 33, 12, 33, 14,
	33, 295, 11, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 314,
	10, 35, 12, 35, 14, 35, 317, 11, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 2, 3, 26, 37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 2, 3, 3, 2, 14, 15, 2, 326, 2, 72, 3, 2, 2, 2, 4, 75, 3, 2,
	2, 2, 6, 80, 3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 96, 3, 2, 2, 2, 12, 101,
	3, 2, 2, 2, 14, 108, 3, 2, 2, 2, 16, 110, 3, 2, 2, 2, 18, 114, 3, 2, 2,
	2, 20, 121, 3, 2, 2, 2, 22, 130, 3, 2, 2, 2, 24, 156, 3, 2, 2, 2, 26, 169,
	3, 2, 2, 2, 28, 184, 3, 2, 2, 2, 30, 186, 3, 2, 2, 2, 32, 188, 3, 2, 2,
	2, 34, 190, 3, 2, 2, 2, 36, 192, 3, 2, 2, 2, 38, 194, 3, 2, 2, 2, 40, 196,
	3, 2, 2, 2, 42, 203, 3, 2, 2, 2, 44, 214, 3, 2, 2, 2, 46, 224, 3, 2, 2,
	2, 48, 226, 3, 2, 2, 2, 50, 229, 3, 2, 2, 2, 52, 233, 3, 2, 2, 2, 54, 236,
	3, 2, 2, 2, 56, 239, 3, 2, 2, 2, 58, 250, 3, 2, 2, 2, 60, 262, 3, 2, 2,
	2, 62, 274, 3, 2, 2, 2, 64, 286, 3, 2, 2, 2, 66, 298, 3, 2, 2, 2, 68, 309,
	3, 2, 2, 2, 70, 320, 3, 2, 2, 2, 72, 73, 5, 4, 3, 2, 73, 74, 7, 2, 2, 3,
	74, 3, 3, 2, 2, 2, 75, 78, 5, 6, 4, 2, 76, 77, 7, 12, 2, 2, 77, 79, 5,
	6, 4, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 5, 3, 2, 2, 2, 80,
	83, 5, 8, 5, 2, 81, 82, 7, 11, 2, 2, 82, 84, 5, 8, 5, 2, 83, 81, 3, 2,
	2, 2, 83, 84, 3, 2, 2, 2, 84, 7, 3, 2, 2, 2, 85, 87, 7, 13, 2, 2, 86, 85,
	3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 5, 10, 6, 2,
	89, 9, 3, 2, 2, 2, 90, 97, 5, 12, 7, 2, 91, 97, 5, 36, 19, 2, 92, 93, 7,
	44, 2, 2, 93, 94, 5, 4, 3, 2, 94, 95, 7, 45, 2, 2, 95, 97, 3, 2, 2, 2,
	96, 90, 3, 2, 2, 2, 96, 91, 3, 2, 2, 2, 96, 92, 3, 2, 2, 2, 97, 11, 3,
	2, 2, 2, 98, 102, 5, 14, 8, 2, 99, 102, 5, 40, 21, 2, 100, 102, 5, 42,
	22, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 100, 3, 2, 2, 2,
	102, 13, 3, 2, 2, 2, 103, 109, 5, 16, 9, 2, 104, 109, 5, 18, 10, 2, 105,
	109, 5, 20, 11, 2, 106, 109, 5, 22, 12, 2, 107, 109, 5, 24, 13, 2, 108,
	103, 3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106,
	3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 15, 3, 2, 2, 2, 110, 111, 5, 26,
	14, 2, 111, 112, 7, 3, 2, 2, 112, 113, 5, 26, 14, 2, 113, 17, 3, 2, 2,
	2, 114, 116, 5, 30, 16, 2, 115, 117, 7, 13, 2, 2, 116, 115, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 9, 2, 2, 2, 119,
	120, 5, 32, 17, 2, 120, 19, 3, 2, 2, 2, 121, 123, 5, 26, 14, 2, 122, 124,
	7, 13, 2, 2, 123, 122, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 3, 2,
	2, 2, 125, 126, 7, 16, 2, 2, 126, 127, 5, 26, 14, 2, 127, 128, 7, 11, 2,
	2, 128, 129, 5, 26, 14, 2, 129, 21, 3, 2, 2, 2, 130, 132, 5, 30, 16, 2,
	131, 133, 7, 13, 2, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 135, 7, 19, 2, 2, 135, 152, 7, 44, 2, 2, 136, 141,
	5, 32, 17, 2, 137, 138, 7, 50, 2, 2, 138, 140, 5, 32, 17, 2, 139, 137,
	3, 2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2,
	2, 2, 142, 153, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 149, 5, 34, 18,
	2, 145, 146, 7, 50, 2, 2, 146, 148, 5, 34, 18, 2, 147, 145, 3, 2, 2, 2,
	148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150,
	153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 136, 3, 2, 2, 2, 152, 144,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 155, 7, 45, 2, 2, 155, 23, 3, 2,
	2, 2, 156, 157, 5, 30, 16, 2, 157, 159, 7, 17, 2, 2, 158, 160, 7, 13, 2,
	2, 159, 158, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161,
	162, 7, 18, 2, 2, 162, 25, 3, 2, 2, 2, 163, 164, 8, 14, 1, 2, 164, 170,
	5, 28, 15, 2, 165, 166, 7, 44, 2, 2, 166, 167, 5, 26, 14, 2, 167, 168,
	7, 45, 2, 2, 168, 170, 3, 2, 2, 2, 169, 163, 3, 2, 2, 2, 169, 165, 3, 2,
	2, 2, 170, 176, 3, 2, 2, 2, 171, 172, 12, 3, 2, 2, 172, 173, 7, 20, 2,
	2, 173, 175, 5, 26, 14, 4, 174, 171, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2,
	176, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 27, 3, 2, 2, 2, 178, 176,
	3, 2, 2, 2, 179, 185, 5, 30, 16, 2, 180, 185, 5, 32, 17, 2, 181, 185, 5,
	34, 18, 2, 182, 185, 5, 36, 19, 2, 183, 185, 5, 38, 20, 2, 184, 179, 3,
	2, 2, 2, 184, 180, 3, 2, 2, 2, 184, 181, 3, 2, 2, 2, 184, 182, 3, 2, 2,
	2, 184, 183, 3, 2, 2, 2, 185, 29, 3, 2, 2, 2, 186, 187, 7, 32, 2, 2, 187,
	31, 3, 2, 2, 2, 188, 189, 7, 84, 2, 2, 189, 33, 3, 2, 2, 2, 190, 191, 7,
	31, 2, 2, 191, 35, 3, 2, 2, 2, 192, 193, 7, 10, 2, 2, 193, 37, 3, 2, 2,
	2, 194, 195, 7, 71, 2, 2, 195, 39, 3, 2, 2, 2, 196, 197, 7, 21, 2, 2, 197,
	198, 7, 44, 2, 2, 198, 199, 5, 44, 23, 2, 199, 200, 7, 50, 2, 2, 200, 201,
	5, 44, 23, 2, 201, 202, 7, 45, 2, 2, 202, 41, 3, 2, 2, 2, 203, 204, 7,
	22, 2, 2, 204, 205, 7, 44, 2, 2, 205, 206, 5, 44, 23, 2, 206, 207, 7, 50,
	2, 2, 207, 208, 5, 44, 23, 2, 208, 209, 7, 50, 2, 2, 209, 210, 7, 31, 2,
	2, 210, 211, 7, 45, 2, 2, 211, 43, 3, 2, 2, 2, 212, 215, 5, 30, 16, 2,
	213, 215, 5, 46, 24, 2, 214, 212, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215,
	45, 3, 2, 2, 2, 216, 225, 5, 48, 25, 2, 217, 225, 5, 52, 27, 2, 218, 225,
	5, 54, 28, 2, 219, 225, 5, 58, 30, 2, 220, 225, 5, 60, 31, 2, 221, 225,
	5, 62, 32, 2, 222, 225, 5, 64, 33, 2, 223, 225, 5, 66, 34, 2, 224, 216,
	3, 2, 2, 2, 224, 217, 3, 2, 2, 2, 224, 218, 3, 2, 2, 2, 224, 219, 3, 2,
	2, 2, 224, 220, 3, 2, 2, 2, 224, 221, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2,
	224, 223, 3, 2, 2, 2, 225, 47, 3, 2, 2, 2, 226, 227, 7, 23, 2, 2, 227,
	228, 5, 50, 26, 2, 228, 49, 3, 2, 2, 2, 229, 230, 7, 44, 2, 2, 230, 231,
	5, 70, 36, 2, 231, 232, 7, 45, 2, 2, 232, 51, 3, 2, 2, 2, 233, 234, 7,
	24, 2, 2, 234, 235, 5, 68, 35, 2, 235, 53, 3, 2, 2, 2, 236, 237, 7, 25,
	2, 2, 237, 238, 5, 56, 29, 2, 238, 55, 3, 2, 2, 2, 239, 240, 7, 44, 2,
	2, 240, 245, 5, 68, 35, 2, 241, 242, 7, 50, 2, 2, 242, 244, 5, 68, 35,
	2, 243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245,
	246, 3, 2, 2, 2, 246, 248, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 248, 249,
	7, 45, 2, 2, 249, 57, 3, 2, 2, 2, 250, 251, 7, 26, 2, 2, 251, 252, 7, 44,
	2, 2, 252, 257, 5, 50, 26, 2, 253, 254, 7, 50, 2, 2, 254, 256, 5, 50, 26,
	2, 255, 253, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 260, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 261,
	7, 45, 2, 2, 261, 59, 3, 2, 2, 2, 262, 263, 7, 27, 2, 2, 263, 264, 7, 44,
	2, 2, 264, 269, 5, 68, 35, 2, 265, 266, 7, 50, 2, 2, 266, 268, 5, 68, 35,
	2, 267, 265, 3, 2, 2, 2, 268, 271, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 272, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 272, 273,
	7, 45, 2, 2, 273, 61, 3, 2, 2, 2, 274, 275, 7, 28, 2, 2, 275, 276, 7, 44,
	2, 2, 276, 281, 5, 56, 29, 2, 277, 278, 7, 50, 2, 2, 278, 280, 5, 56, 29,
	2, 279, 277, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 281,
	282, 3, 2, 2, 2, 282, 284, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 284, 285,
	7, 45, 2, 2, 285, 63, 3, 2, 2, 2, 286, 287, 7, 29, 2, 2, 287, 288, 7, 44,
	2, 2, 288, 293, 5, 46, 24, 2, 289, 290, 7, 50, 2, 2, 290, 292, 5, 46, 24,
	2, 291, 289, 3, 2, 2, 2, 292, 295, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 293,
	294, 3, 2, 2, 2, 294, 296, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 296, 297,
	7, 45, 2, 2, 297, 65, 3, 2, 2, 2, 298, 299, 7, 30, 2, 2, 299, 300, 7, 44,
	2, 2, 300, 301, 7, 31, 2, 2, 301, 302, 7, 50, 2, 2, 302, 303, 7, 31, 2,
	2, 303, 304, 7, 50, 2, 2, 304, 305, 7, 31, 2, 2, 305, 306, 7, 50, 2, 2,
	306, 307, 7, 31, 2, 2, 307, 308, 7, 45, 2, 2, 308, 67, 3, 2, 2, 2, 309,
	310, 7, 44, 2, 2, 310, 315, 5, 70, 36, 2, 311, 312, 7, 50, 2, 2, 312, 314,
	5, 70, 36, 2, 313, 311, 3, 2, 2, 2, 314, 317, 3, 2, 2, 2, 315, 313, 3,
	2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2, 2, 317, 315, 3, 2, 2,
	2, 318, 319, 7, 45, 2, 2, 319, 69, 3, 2, 2, 2, 320, 321, 7, 31, 2, 2, 321,
	322, 7, 31, 2, 2, 322, 71, 3, 2, 2, 2, 26, 78, 83, 86, 96, 101, 108, 116,
	123, 132, 141, 149, 152, 159, 169, 176, 184, 214, 224, 245, 257, 269, 281,
	293, 315,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "'#'", "'$'", "'_'", "'\"'", "'%'", "'&'", "", "'('", "')'", "'['",
	"']'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "'^'", "'||'", "':'",
	"';'", "'?'", "'|'", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "''''",
}
var symbolicNames = []string{
	"", "ComparisonOperator", "LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral",
	"AND", "OR", "NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "IN", "ArithmeticOperator",
	"SpatialOperator", "DistanceOperator", "POINT", "LINESTRING", "POLYGON",
	"MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON", "GEOMETRYCOLLECTION",
	"ENVELOPE", "NumericLiteral", "Identifier", "IdentifierStart", "IdentifierPart",
	"ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR", "UNDERSCORE", "DOUBLEQUOTE", "PERCENT",
	"AMPERSAND", "QUOTE", "LEFTPAREN", "RIGHTPAREN", "LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET",
	"ASTERISK", "PLUS", "COMMA", "MINUS", "PERIOD", "SOLIDUS", "CARET", "CONCAT",
	"COLON", "SEMICOLON", "QUESTIONMARK", "VERTICALBAR", "BIT", "HEXIT", "UnsignedNumericLiteral",
	"SignedNumericLiteral", "ExactNumericLiteral", "ApproximateNumericLiteral",
	"Mantissa", "Exponent", "SignedInteger", "UnsignedInteger", "Sign", "TemporalLiteral",
	"Instant", "FullDate", "DateYear", "DateMonth", "DateDay", "UtcTime", "TimeZoneOffset",
	"TimeHour", "TimeMinute", "TimeSecond", "NOW", "WS", "CharacterStringLiteral",
	"QuotedQuote",
}

var ruleNames = []string{
	"cqlFilter", "booleanExpression", "booleanTerm", "booleanFactor", "booleanPrimary",
	"predicate", "comparisonPredicate", "binaryComparisonPredicate", "isLikePredicate",
	"isBetweenPredicate", "isInListPredicate", "isNullPredicate", "scalarExpression",
	"scalarValue", "propertyName", "characterLiteral", "numericLiteral", "booleanLiteral",
	"temporalLiteral", "spatialPredicate", "distancePredicate", "geomExpression",
	"geomLiteral", "point", "pointList", "linestring", "polygon", "polygonDef",
	"multiPoint", "multiLinestring", "multiPolygon", "geometryCollection",
	"envelope", "coordList", "coordinate",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CQLParser struct {
	*antlr.BaseParser
}

func NewCQLParser(input antlr.TokenStream) *CQLParser {
	this := new(CQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CQLParser.g4"

	return this
}

// CQLParser tokens.
const (
	CQLParserEOF                       = antlr.TokenEOF
	CQLParserComparisonOperator        = 1
	CQLParserLT                        = 2
	CQLParserEQ                        = 3
	CQLParserGT                        = 4
	CQLParserNEQ                       = 5
	CQLParserGTEQ                      = 6
	CQLParserLTEQ                      = 7
	CQLParserBooleanLiteral            = 8
	CQLParserAND                       = 9
	CQLParserOR                        = 10
	CQLParserNOT                       = 11
	CQLParserLIKE                      = 12
	CQLParserILIKE                     = 13
	CQLParserBETWEEN                   = 14
	CQLParserIS                        = 15
	CQLParserNULL                      = 16
	CQLParserIN                        = 17
	CQLParserArithmeticOperator        = 18
	CQLParserSpatialOperator           = 19
	CQLParserDistanceOperator          = 20
	CQLParserPOINT                     = 21
	CQLParserLINESTRING                = 22
	CQLParserPOLYGON                   = 23
	CQLParserMULTIPOINT                = 24
	CQLParserMULTILINESTRING           = 25
	CQLParserMULTIPOLYGON              = 26
	CQLParserGEOMETRYCOLLECTION        = 27
	CQLParserENVELOPE                  = 28
	CQLParserNumericLiteral            = 29
	CQLParserIdentifier                = 30
	CQLParserIdentifierStart           = 31
	CQLParserIdentifierPart            = 32
	CQLParserALPHA                     = 33
	CQLParserDIGIT                     = 34
	CQLParserOCTOTHORP                 = 35
	CQLParserDOLLAR                    = 36
	CQLParserUNDERSCORE                = 37
	CQLParserDOUBLEQUOTE               = 38
	CQLParserPERCENT                   = 39
	CQLParserAMPERSAND                 = 40
	CQLParserQUOTE                     = 41
	CQLParserLEFTPAREN                 = 42
	CQLParserRIGHTPAREN                = 43
	CQLParserLEFTSQUAREBRACKET         = 44
	CQLParserRIGHTSQUAREBRACKET        = 45
	CQLParserASTERISK                  = 46
	CQLParserPLUS                      = 47
	CQLParserCOMMA                     = 48
	CQLParserMINUS                     = 49
	CQLParserPERIOD                    = 50
	CQLParserSOLIDUS                   = 51
	CQLParserCARET                     = 52
	CQLParserCONCAT                    = 53
	CQLParserCOLON                     = 54
	CQLParserSEMICOLON                 = 55
	CQLParserQUESTIONMARK              = 56
	CQLParserVERTICALBAR               = 57
	CQLParserBIT                       = 58
	CQLParserHEXIT                     = 59
	CQLParserUnsignedNumericLiteral    = 60
	CQLParserSignedNumericLiteral      = 61
	CQLParserExactNumericLiteral       = 62
	CQLParserApproximateNumericLiteral = 63
	CQLParserMantissa                  = 64
	CQLParserExponent                  = 65
	CQLParserSignedInteger             = 66
	CQLParserUnsignedInteger           = 67
	CQLParserSign                      = 68
	CQLParserTemporalLiteral           = 69
	CQLParserInstant                   = 70
	CQLParserFullDate                  = 71
	CQLParserDateYear                  = 72
	CQLParserDateMonth                 = 73
	CQLParserDateDay                   = 74
	CQLParserUtcTime                   = 75
	CQLParserTimeZoneOffset            = 76
	CQLParserTimeHour                  = 77
	CQLParserTimeMinute                = 78
	CQLParserTimeSecond                = 79
	CQLParserNOW                       = 80
	CQLParserWS                        = 81
	CQLParserCharacterStringLiteral    = 82
	CQLParserQuotedQuote               = 83
)

// CQLParser rules.
const (
	CQLParserRULE_cqlFilter                 = 0
	CQLParserRULE_booleanExpression         = 1
	CQLParserRULE_booleanTerm               = 2
	CQLParserRULE_booleanFactor             = 3
	CQLParserRULE_booleanPrimary            = 4
	CQLParserRULE_predicate                 = 5
	CQLParserRULE_comparisonPredicate       = 6
	CQLParserRULE_binaryComparisonPredicate = 7
	CQLParserRULE_isLikePredicate           = 8
	CQLParserRULE_isBetweenPredicate        = 9
	CQLParserRULE_isInListPredicate         = 10
	CQLParserRULE_isNullPredicate           = 11
	CQLParserRULE_scalarExpression          = 12
	CQLParserRULE_scalarValue               = 13
	CQLParserRULE_propertyName              = 14
	CQLParserRULE_characterLiteral          = 15
	CQLParserRULE_numericLiteral            = 16
	CQLParserRULE_booleanLiteral            = 17
	CQLParserRULE_temporalLiteral           = 18
	CQLParserRULE_spatialPredicate          = 19
	CQLParserRULE_distancePredicate         = 20
	CQLParserRULE_geomExpression            = 21
	CQLParserRULE_geomLiteral               = 22
	CQLParserRULE_point                     = 23
	CQLParserRULE_pointList                 = 24
	CQLParserRULE_linestring                = 25
	CQLParserRULE_polygon                   = 26
	CQLParserRULE_polygonDef                = 27
	CQLParserRULE_multiPoint                = 28
	CQLParserRULE_multiLinestring           = 29
	CQLParserRULE_multiPolygon              = 30
	CQLParserRULE_geometryCollection        = 31
	CQLParserRULE_envelope                  = 32
	CQLParserRULE_coordList                 = 33
	CQLParserRULE_coordinate                = 34
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCqlFilterContext() *CqlFilterContext {
	var p = new(CqlFilterContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_cqlFilter
	return p
}

func (*CqlFilterContext) IsCqlFilterContext() {}

func NewCqlFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlFilterContext {
	var p = new(CqlFilterContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_cqlFilter

	return p
}

func (s *CqlFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlFilterContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *CqlFilterContext) EOF() antlr.TerminalNode {
	return s.GetToken(CQLParserEOF, 0)
}

func (s *CqlFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCqlFilter(s)
	}
}

func (s *CqlFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCqlFilter(s)
	}
}

func (p *CQLParser) CqlFilter() (localctx ICqlFilterContext) {
	localctx = NewCqlFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CQLParserRULE_cqlFilter)

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
		p.SetState(70)
		p.BooleanExpression()
	}
	{
		p.SetState(71)
		p.Match(CQLParserEOF)
	}

	return localctx
}

// IBooleanExpressionContext is an interface to support dynamic dispatch.
type IBooleanExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanExpressionContext differentiates from other interfaces.
	IsBooleanExpressionContext()
}

type BooleanExpressionContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanExpressionContext() *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanExpression
	return p
}

func (*BooleanExpressionContext) IsBooleanExpressionContext() {}

func NewBooleanExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanExpression

	return p
}

func (s *BooleanExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExpressionContext) AllBooleanTerm() []IBooleanTermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem())
	var tst = make([]IBooleanTermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBooleanTermContext)
		}
	}

	return tst
}

func (s *BooleanExpressionContext) BooleanTerm(i int) IBooleanTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanTermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBooleanTermContext)
}

func (s *BooleanExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(CQLParserOR, 0)
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanExpression(s)
	}
}

func (s *BooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanExpression(s)
	}
}

func (p *CQLParser) BooleanExpression() (localctx IBooleanExpressionContext) {
	localctx = NewBooleanExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CQLParserRULE_booleanExpression)
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
		p.SetState(73)
		p.BooleanTerm()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserOR {
		{
			p.SetState(74)
			p.Match(CQLParserOR)
		}
		{
			p.SetState(75)
			p.BooleanTerm()
		}

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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanTermContext() *BooleanTermContext {
	var p = new(BooleanTermContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanTerm
	return p
}

func (*BooleanTermContext) IsBooleanTermContext() {}

func NewBooleanTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanTermContext {
	var p = new(BooleanTermContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanTerm

	return p
}

func (s *BooleanTermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanTermContext) AllBooleanFactor() []IBooleanFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBooleanFactorContext)(nil)).Elem())
	var tst = make([]IBooleanFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBooleanFactorContext)
		}
	}

	return tst
}

func (s *BooleanTermContext) BooleanFactor(i int) IBooleanFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBooleanFactorContext)
}

func (s *BooleanTermContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *BooleanTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanTerm(s)
	}
}

func (s *BooleanTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanTerm(s)
	}
}

func (p *CQLParser) BooleanTerm() (localctx IBooleanTermContext) {
	localctx = NewBooleanTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CQLParserRULE_booleanTerm)
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
		p.SetState(78)
		p.BooleanFactor()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserAND {
		{
			p.SetState(79)
			p.Match(CQLParserAND)
		}
		{
			p.SetState(80)
			p.BooleanFactor()
		}

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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanFactorContext() *BooleanFactorContext {
	var p = new(BooleanFactorContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanFactor
	return p
}

func (*BooleanFactorContext) IsBooleanFactorContext() {}

func NewBooleanFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFactorContext {
	var p = new(BooleanFactorContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanFactor

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
	return s.GetToken(CQLParserNOT, 0)
}

func (s *BooleanFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanFactor(s)
	}
}

func (s *BooleanFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanFactor(s)
	}
}

func (p *CQLParser) BooleanFactor() (localctx IBooleanFactorContext) {
	localctx = NewBooleanFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CQLParserRULE_booleanFactor)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(83)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(86)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanPrimaryContext() *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanPrimary
	return p
}

func (*BooleanPrimaryContext) IsBooleanPrimaryContext() {}

func NewBooleanPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanPrimaryContext {
	var p = new(BooleanPrimaryContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanPrimary

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

func (s *BooleanPrimaryContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *BooleanPrimaryContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *BooleanPrimaryContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *BooleanPrimaryContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *BooleanPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanPrimary(s)
	}
}

func (s *BooleanPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanPrimary(s)
	}
}

func (p *CQLParser) BooleanPrimary() (localctx IBooleanPrimaryContext) {
	localctx = NewBooleanPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CQLParserRULE_booleanPrimary)

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

	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Predicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.BooleanLiteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(CQLParserLEFTPAREN)
		}
		{
			p.SetState(91)
			p.BooleanExpression()
		}
		{
			p.SetState(92)
			p.Match(CQLParserRIGHTPAREN)
		}

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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) ComparisonPredicate() IComparisonPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonPredicateContext)
}

func (s *PredicateContext) SpatialPredicate() ISpatialPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpatialPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpatialPredicateContext)
}

func (s *PredicateContext) DistancePredicate() IDistancePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDistancePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDistancePredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *CQLParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CQLParserRULE_predicate)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserBooleanLiteral, CQLParserNumericLiteral, CQLParserIdentifier, CQLParserLEFTPAREN, CQLParserTemporalLiteral, CQLParserCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.ComparisonPredicate()
		}

	case CQLParserSpatialOperator:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.SpatialPredicate()
		}

	case CQLParserDistanceOperator:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.DistancePredicate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonPredicateContext is an interface to support dynamic dispatch.
type IComparisonPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonPredicateContext differentiates from other interfaces.
	IsComparisonPredicateContext()
}

type ComparisonPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyComparisonPredicateContext() *ComparisonPredicateContext {
	var p = new(ComparisonPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_comparisonPredicate
	return p
}

func (*ComparisonPredicateContext) IsComparisonPredicateContext() {}

func NewComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonPredicateContext {
	var p = new(ComparisonPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_comparisonPredicate

	return p
}

func (s *ComparisonPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonPredicateContext) BinaryComparisonPredicate() IBinaryComparisonPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryComparisonPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryComparisonPredicateContext)
}

func (s *ComparisonPredicateContext) IsLikePredicate() IIsLikePredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsLikePredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsLikePredicateContext)
}

func (s *ComparisonPredicateContext) IsBetweenPredicate() IIsBetweenPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsBetweenPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsBetweenPredicateContext)
}

func (s *ComparisonPredicateContext) IsInListPredicate() IIsInListPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsInListPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsInListPredicateContext)
}

func (s *ComparisonPredicateContext) IsNullPredicate() IIsNullPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsNullPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsNullPredicateContext)
}

func (s *ComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterComparisonPredicate(s)
	}
}

func (s *ComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitComparisonPredicate(s)
	}
}

func (p *CQLParser) ComparisonPredicate() (localctx IComparisonPredicateContext) {
	localctx = NewComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CQLParserRULE_comparisonPredicate)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.BinaryComparisonPredicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.IsLikePredicate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.IsBetweenPredicate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.IsInListPredicate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.IsNullPredicate()
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBinaryComparisonPredicateContext() *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate
	return p
}

func (*BinaryComparisonPredicateContext) IsBinaryComparisonPredicateContext() {}

func NewBinaryComparisonPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryComparisonPredicateContext {
	var p = new(BinaryComparisonPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_binaryComparisonPredicate

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
	return s.GetToken(CQLParserComparisonOperator, 0)
}

func (s *BinaryComparisonPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparisonPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryComparisonPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBinaryComparisonPredicate(s)
	}
}

func (s *BinaryComparisonPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBinaryComparisonPredicate(s)
	}
}

func (p *CQLParser) BinaryComparisonPredicate() (localctx IBinaryComparisonPredicateContext) {
	localctx = NewBinaryComparisonPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CQLParserRULE_binaryComparisonPredicate)

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
		p.SetState(108)
		p.scalarExpression(0)
	}
	{
		p.SetState(109)
		p.Match(CQLParserComparisonOperator)
	}
	{
		p.SetState(110)
		p.scalarExpression(0)
	}

	return localctx
}

// IIsLikePredicateContext is an interface to support dynamic dispatch.
type IIsLikePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsLikePredicateContext differentiates from other interfaces.
	IsIsLikePredicateContext()
}

type IsLikePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsLikePredicateContext() *IsLikePredicateContext {
	var p = new(IsLikePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_isLikePredicate
	return p
}

func (*IsLikePredicateContext) IsIsLikePredicateContext() {}

func NewIsLikePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsLikePredicateContext {
	var p = new(IsLikePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_isLikePredicate

	return p
}

func (s *IsLikePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsLikePredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *IsLikePredicateContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *IsLikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserLIKE, 0)
}

func (s *IsLikePredicateContext) ILIKE() antlr.TerminalNode {
	return s.GetToken(CQLParserILIKE, 0)
}

func (s *IsLikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsLikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsLikePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsLikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsLikePredicate(s)
	}
}

func (s *IsLikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsLikePredicate(s)
	}
}

func (p *CQLParser) IsLikePredicate() (localctx IIsLikePredicateContext) {
	localctx = NewIsLikePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CQLParserRULE_isLikePredicate)
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
		p.SetState(112)
		p.PropertyName()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(113)
			p.Match(CQLParserNOT)
		}

	}
	p.SetState(116)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLParserLIKE || _la == CQLParserILIKE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(117)
		p.CharacterLiteral()
	}

	return localctx
}

// IIsBetweenPredicateContext is an interface to support dynamic dispatch.
type IIsBetweenPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsBetweenPredicateContext differentiates from other interfaces.
	IsIsBetweenPredicateContext()
}

type IsBetweenPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsBetweenPredicateContext() *IsBetweenPredicateContext {
	var p = new(IsBetweenPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_isBetweenPredicate
	return p
}

func (*IsBetweenPredicateContext) IsIsBetweenPredicateContext() {}

func NewIsBetweenPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsBetweenPredicateContext {
	var p = new(IsBetweenPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_isBetweenPredicate

	return p
}

func (s *IsBetweenPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsBetweenPredicateContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *IsBetweenPredicateContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *IsBetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(CQLParserBETWEEN, 0)
}

func (s *IsBetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(CQLParserAND, 0)
}

func (s *IsBetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsBetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsBetweenPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsBetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsBetweenPredicate(s)
	}
}

func (s *IsBetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsBetweenPredicate(s)
	}
}

func (p *CQLParser) IsBetweenPredicate() (localctx IIsBetweenPredicateContext) {
	localctx = NewIsBetweenPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CQLParserRULE_isBetweenPredicate)
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
		p.SetState(119)
		p.scalarExpression(0)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(120)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(123)
		p.Match(CQLParserBETWEEN)
	}
	{
		p.SetState(124)
		p.scalarExpression(0)
	}
	{
		p.SetState(125)
		p.Match(CQLParserAND)
	}
	{
		p.SetState(126)
		p.scalarExpression(0)
	}

	return localctx
}

// IIsInListPredicateContext is an interface to support dynamic dispatch.
type IIsInListPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsInListPredicateContext differentiates from other interfaces.
	IsIsInListPredicateContext()
}

type IsInListPredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsInListPredicateContext() *IsInListPredicateContext {
	var p = new(IsInListPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_isInListPredicate
	return p
}

func (*IsInListPredicateContext) IsIsInListPredicateContext() {}

func NewIsInListPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsInListPredicateContext {
	var p = new(IsInListPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_isInListPredicate

	return p
}

func (s *IsInListPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *IsInListPredicateContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *IsInListPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(CQLParserIN, 0)
}

func (s *IsInListPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *IsInListPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *IsInListPredicateContext) AllCharacterLiteral() []ICharacterLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem())
	var tst = make([]ICharacterLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharacterLiteralContext)
		}
	}

	return tst
}

func (s *IsInListPredicateContext) CharacterLiteral(i int) ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *IsInListPredicateContext) AllNumericLiteral() []INumericLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem())
	var tst = make([]INumericLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralContext)
		}
	}

	return tst
}

func (s *IsInListPredicateContext) NumericLiteral(i int) INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *IsInListPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsInListPredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *IsInListPredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *IsInListPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsInListPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsInListPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsInListPredicate(s)
	}
}

func (s *IsInListPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsInListPredicate(s)
	}
}

func (p *CQLParser) IsInListPredicate() (localctx IIsInListPredicateContext) {
	localctx = NewIsInListPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CQLParserRULE_isInListPredicate)
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
		p.SetState(128)
		p.PropertyName()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(129)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(132)
		p.Match(CQLParserIN)
	}
	{
		p.SetState(133)
		p.Match(CQLParserLEFTPAREN)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserCharacterStringLiteral:
		{
			p.SetState(134)
			p.CharacterLiteral()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(135)
				p.Match(CQLParserCOMMA)
			}
			{
				p.SetState(136)
				p.CharacterLiteral()
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case CQLParserNumericLiteral:
		{
			p.SetState(142)
			p.NumericLiteral()
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLParserCOMMA {
			{
				p.SetState(143)
				p.Match(CQLParserCOMMA)
			}
			{
				p.SetState(144)
				p.NumericLiteral()
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(152)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyIsNullPredicateContext() *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_isNullPredicate
	return p
}

func (*IsNullPredicateContext) IsIsNullPredicateContext() {}

func NewIsNullPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_isNullPredicate

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
	return s.GetToken(CQLParserIS, 0)
}

func (s *IsNullPredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(CQLParserNULL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(CQLParserNOT, 0)
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (p *CQLParser) IsNullPredicate() (localctx IIsNullPredicateContext) {
	localctx = NewIsNullPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CQLParserRULE_isNullPredicate)
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
		p.SetState(154)
		p.PropertyName()
	}
	{
		p.SetState(155)
		p.Match(CQLParserIS)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLParserNOT {
		{
			p.SetState(156)
			p.Match(CQLParserNOT)
		}

	}
	{
		p.SetState(159)
		p.Match(CQLParserNULL)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarExpressionContext() *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_scalarExpression
	return p
}

func (*ScalarExpressionContext) IsScalarExpressionContext() {}

func NewScalarExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarExpressionContext {
	var p = new(ScalarExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarExpression

	return p
}

func (s *ScalarExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarExpressionContext) ScalarValue() IScalarValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarValueContext)
}

func (s *ScalarExpressionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *ScalarExpressionContext) AllScalarExpression() []IScalarExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem())
	var tst = make([]IScalarExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScalarExpressionContext)
		}
	}

	return tst
}

func (s *ScalarExpressionContext) ScalarExpression(i int) IScalarExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScalarExpressionContext)
}

func (s *ScalarExpressionContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *ScalarExpressionContext) ArithmeticOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserArithmeticOperator, 0)
}

func (s *ScalarExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarExpression(s)
	}
}

func (s *ScalarExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarExpression(s)
	}
}

func (p *CQLParser) ScalarExpression() (localctx IScalarExpressionContext) {
	return p.scalarExpression(0)
}

func (p *CQLParser) scalarExpression(_p int) (localctx IScalarExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewScalarExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IScalarExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, CQLParserRULE_scalarExpression, _p)

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
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserBooleanLiteral, CQLParserNumericLiteral, CQLParserIdentifier, CQLParserTemporalLiteral, CQLParserCharacterStringLiteral:
		{
			p.SetState(162)
			p.ScalarValue()
		}

	case CQLParserLEFTPAREN:
		{
			p.SetState(163)
			p.Match(CQLParserLEFTPAREN)
		}
		{
			p.SetState(164)
			p.scalarExpression(0)
		}
		{
			p.SetState(165)
			p.Match(CQLParserRIGHTPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewScalarExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CQLParserRULE_scalarExpression)
			p.SetState(169)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(170)
				p.Match(CQLParserArithmeticOperator)
			}
			{
				p.SetState(171)
				p.scalarExpression(2)
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IScalarValueContext is an interface to support dynamic dispatch.
type IScalarValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarValueContext differentiates from other interfaces.
	IsScalarValueContext()
}

type ScalarValueContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyScalarValueContext() *ScalarValueContext {
	var p = new(ScalarValueContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_scalarValue
	return p
}

func (*ScalarValueContext) IsScalarValueContext() {}

func NewScalarValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarValueContext {
	var p = new(ScalarValueContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_scalarValue

	return p
}

func (s *ScalarValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarValueContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *ScalarValueContext) CharacterLiteral() ICharacterLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharacterLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICharacterLiteralContext)
}

func (s *ScalarValueContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *ScalarValueContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ScalarValueContext) TemporalLiteral() ITemporalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporalLiteralContext)
}

func (s *ScalarValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterScalarValue(s)
	}
}

func (s *ScalarValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitScalarValue(s)
	}
}

func (p *CQLParser) ScalarValue() (localctx IScalarValueContext) {
	localctx = NewScalarValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CQLParserRULE_scalarValue)

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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.PropertyName()
		}

	case CQLParserCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)
			p.CharacterLiteral()
		}

	case CQLParserNumericLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(179)
			p.NumericLiteral()
		}

	case CQLParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(180)
			p.BooleanLiteral()
		}

	case CQLParserTemporalLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(181)
			p.TemporalLiteral()
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CQLParserIdentifier, 0)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *CQLParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CQLParserRULE_propertyName)

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
		p.SetState(184)
		p.Match(CQLParserIdentifier)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCharacterLiteralContext() *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_characterLiteral
	return p
}

func (*CharacterLiteralContext) IsCharacterLiteralContext() {}

func NewCharacterLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharacterLiteralContext {
	var p = new(CharacterLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_characterLiteral

	return p
}

func (s *CharacterLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *CharacterLiteralContext) CharacterStringLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserCharacterStringLiteral, 0)
}

func (s *CharacterLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharacterLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharacterLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCharacterLiteral(s)
	}
}

func (s *CharacterLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCharacterLiteral(s)
	}
}

func (p *CQLParser) CharacterLiteral() (localctx ICharacterLiteralContext) {
	localctx = NewCharacterLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CQLParserRULE_characterLiteral)

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
		p.SetState(186)
		p.Match(CQLParserCharacterStringLiteral)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *CQLParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLParserRULE_numericLiteral)

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
		p.SetState(188)
		p.Match(CQLParserNumericLiteral)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *CQLParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CQLParserRULE_booleanLiteral)

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
		p.Match(CQLParserBooleanLiteral)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyTemporalLiteralContext() *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_temporalLiteral
	return p
}

func (*TemporalLiteralContext) IsTemporalLiteralContext() {}

func NewTemporalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalLiteralContext {
	var p = new(TemporalLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_temporalLiteral

	return p
}

func (s *TemporalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalLiteralContext) TemporalLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserTemporalLiteral, 0)
}

func (s *TemporalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterTemporalLiteral(s)
	}
}

func (s *TemporalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitTemporalLiteral(s)
	}
}

func (p *CQLParser) TemporalLiteral() (localctx ITemporalLiteralContext) {
	localctx = NewTemporalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CQLParserRULE_temporalLiteral)

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
		p.SetState(192)
		p.Match(CQLParserTemporalLiteral)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptySpatialPredicateContext() *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_spatialPredicate
	return p
}

func (*SpatialPredicateContext) IsSpatialPredicateContext() {}

func NewSpatialPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpatialPredicateContext {
	var p = new(SpatialPredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_spatialPredicate

	return p
}

func (s *SpatialPredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpatialPredicateContext) SpatialOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserSpatialOperator, 0)
}

func (s *SpatialPredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
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
	return s.GetToken(CQLParserCOMMA, 0)
}

func (s *SpatialPredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *SpatialPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpatialPredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpatialPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterSpatialPredicate(s)
	}
}

func (s *SpatialPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitSpatialPredicate(s)
	}
}

func (p *CQLParser) SpatialPredicate() (localctx ISpatialPredicateContext) {
	localctx = NewSpatialPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CQLParserRULE_spatialPredicate)

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
		p.SetState(194)
		p.Match(CQLParserSpatialOperator)
	}
	{
		p.SetState(195)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(196)
		p.GeomExpression()
	}
	{
		p.SetState(197)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(198)
		p.GeomExpression()
	}
	{
		p.SetState(199)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// IDistancePredicateContext is an interface to support dynamic dispatch.
type IDistancePredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDistancePredicateContext differentiates from other interfaces.
	IsDistancePredicateContext()
}

type DistancePredicateContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyDistancePredicateContext() *DistancePredicateContext {
	var p = new(DistancePredicateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_distancePredicate
	return p
}

func (*DistancePredicateContext) IsDistancePredicateContext() {}

func NewDistancePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistancePredicateContext {
	var p = new(DistancePredicateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_distancePredicate

	return p
}

func (s *DistancePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DistancePredicateContext) DistanceOperator() antlr.TerminalNode {
	return s.GetToken(CQLParserDistanceOperator, 0)
}

func (s *DistancePredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *DistancePredicateContext) AllGeomExpression() []IGeomExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem())
	var tst = make([]IGeomExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeomExpressionContext)
		}
	}

	return tst
}

func (s *DistancePredicateContext) GeomExpression(i int) IGeomExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeomExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeomExpressionContext)
}

func (s *DistancePredicateContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *DistancePredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *DistancePredicateContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, 0)
}

func (s *DistancePredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *DistancePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistancePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistancePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterDistancePredicate(s)
	}
}

func (s *DistancePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitDistancePredicate(s)
	}
}

func (p *CQLParser) DistancePredicate() (localctx IDistancePredicateContext) {
	localctx = NewDistancePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CQLParserRULE_distancePredicate)

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
		p.SetState(201)
		p.Match(CQLParserDistanceOperator)
	}
	{
		p.SetState(202)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(203)
		p.GeomExpression()
	}
	{
		p.SetState(204)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(205)
		p.GeomExpression()
	}
	{
		p.SetState(206)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(207)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(208)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomExpressionContext() *GeomExpressionContext {
	var p = new(GeomExpressionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geomExpression
	return p
}

func (*GeomExpressionContext) IsGeomExpressionContext() {}

func NewGeomExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomExpressionContext {
	var p = new(GeomExpressionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomExpression

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
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomExpression(s)
	}
}

func (s *GeomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomExpression(s)
	}
}

func (p *CQLParser) GeomExpression() (localctx IGeomExpressionContext) {
	localctx = NewGeomExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQLParserRULE_geomExpression)

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.PropertyName()
		}

	case CQLParserPOINT, CQLParserLINESTRING, CQLParserPOLYGON, CQLParserMULTIPOINT, CQLParserMULTILINESTRING, CQLParserMULTIPOLYGON, CQLParserGEOMETRYCOLLECTION, CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeomLiteralContext() *GeomLiteralContext {
	var p = new(GeomLiteralContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geomLiteral
	return p
}

func (*GeomLiteralContext) IsGeomLiteralContext() {}

func NewGeomLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeomLiteralContext {
	var p = new(GeomLiteralContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geomLiteral

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
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeomLiteral(s)
	}
}

func (s *GeomLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeomLiteral(s)
	}
}

func (p *CQLParser) GeomLiteral() (localctx IGeomLiteralContext) {
	localctx = NewGeomLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CQLParserRULE_geomLiteral)

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

	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLParserPOINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Point()
		}

	case CQLParserLINESTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Linestring()
		}

	case CQLParserPOLYGON:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Polygon()
		}

	case CQLParserMULTIPOINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.MultiPoint()
		}

	case CQLParserMULTILINESTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.MultiLinestring()
		}

	case CQLParserMULTIPOLYGON:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(219)
			p.MultiPolygon()
		}

	case CQLParserGEOMETRYCOLLECTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(220)
			p.GeometryCollection()
		}

	case CQLParserENVELOPE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(221)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_point
	return p
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) POINT() antlr.TerminalNode {
	return s.GetToken(CQLParserPOINT, 0)
}

func (s *PointContext) PointList() IPointListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *CQLParser) Point() (localctx IPointContext) {
	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CQLParserRULE_point)

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
		p.SetState(224)
		p.Match(CQLParserPOINT)
	}
	{
		p.SetState(225)
		p.PointList()
	}

	return localctx
}

// IPointListContext is an interface to support dynamic dispatch.
type IPointListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointListContext differentiates from other interfaces.
	IsPointListContext()
}

type PointListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPointListContext() *PointListContext {
	var p = new(PointListContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_pointList
	return p
}

func (*PointListContext) IsPointListContext() {}

func NewPointListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointListContext {
	var p = new(PointListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_pointList

	return p
}

func (s *PointListContext) GetParser() antlr.Parser { return s.parser }

func (s *PointListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PointListContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *PointListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PointListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPointList(s)
	}
}

func (s *PointListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPointList(s)
	}
}

func (p *CQLParser) PointList() (localctx IPointListContext) {
	localctx = NewPointListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CQLParserRULE_pointList)

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
		p.SetState(227)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(228)
		p.Coordinate()
	}
	{
		p.SetState(229)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyLinestringContext() *LinestringContext {
	var p = new(LinestringContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_linestring
	return p
}

func (*LinestringContext) IsLinestringContext() {}

func NewLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinestringContext {
	var p = new(LinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_linestring

	return p
}

func (s *LinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *LinestringContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserLINESTRING, 0)
}

func (s *LinestringContext) CoordList() ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *LinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterLinestring(s)
	}
}

func (s *LinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitLinestring(s)
	}
}

func (p *CQLParser) Linestring() (localctx ILinestringContext) {
	localctx = NewLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CQLParserRULE_linestring)

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
		p.SetState(231)
		p.Match(CQLParserLINESTRING)
	}
	{
		p.SetState(232)
		p.CoordList()
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_polygon
	return p
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserPOLYGON, 0)
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
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *CQLParser) Polygon() (localctx IPolygonContext) {
	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CQLParserRULE_polygon)

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
		p.SetState(234)
		p.Match(CQLParserPOLYGON)
	}
	{
		p.SetState(235)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyPolygonDefContext() *PolygonDefContext {
	var p = new(PolygonDefContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_polygonDef
	return p
}

func (*PolygonDefContext) IsPolygonDefContext() {}

func NewPolygonDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonDefContext {
	var p = new(PolygonDefContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_polygonDef

	return p
}

func (s *PolygonDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonDefContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *PolygonDefContext) AllCoordList() []ICoordListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordListContext)(nil)).Elem())
	var tst = make([]ICoordListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordListContext)
		}
	}

	return tst
}

func (s *PolygonDefContext) CoordList(i int) ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *PolygonDefContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *PolygonDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *PolygonDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *PolygonDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterPolygonDef(s)
	}
}

func (s *PolygonDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitPolygonDef(s)
	}
}

func (p *CQLParser) PolygonDef() (localctx IPolygonDefContext) {
	localctx = NewPolygonDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CQLParserRULE_polygonDef)
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
		p.SetState(237)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(238)
		p.CoordList()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(239)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(240)
			p.CoordList()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(246)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPointContext() *MultiPointContext {
	var p = new(MultiPointContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiPoint
	return p
}

func (*MultiPointContext) IsMultiPointContext() {}

func NewMultiPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointContext {
	var p = new(MultiPointContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPoint

	return p
}

func (s *MultiPointContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOINT, 0)
}

func (s *MultiPointContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiPointContext) AllPointList() []IPointListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointListContext)(nil)).Elem())
	var tst = make([]IPointListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointListContext)
		}
	}

	return tst
}

func (s *MultiPointContext) PointList(i int) IPointListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointListContext)
}

func (s *MultiPointContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPointContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPointContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPoint(s)
	}
}

func (s *MultiPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPoint(s)
	}
}

func (p *CQLParser) MultiPoint() (localctx IMultiPointContext) {
	localctx = NewMultiPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CQLParserRULE_multiPoint)
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
		p.SetState(248)
		p.Match(CQLParserMULTIPOINT)
	}
	{
		p.SetState(249)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(250)
		p.PointList()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(251)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(252)
			p.PointList()
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(258)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiLinestringContext() *MultiLinestringContext {
	var p = new(MultiLinestringContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiLinestring
	return p
}

func (*MultiLinestringContext) IsMultiLinestringContext() {}

func NewMultiLinestringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLinestringContext {
	var p = new(MultiLinestringContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiLinestring

	return p
}

func (s *MultiLinestringContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLinestringContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTILINESTRING, 0)
}

func (s *MultiLinestringContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *MultiLinestringContext) AllCoordList() []ICoordListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordListContext)(nil)).Elem())
	var tst = make([]ICoordListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordListContext)
		}
	}

	return tst
}

func (s *MultiLinestringContext) CoordList(i int) ICoordListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordListContext)
}

func (s *MultiLinestringContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiLinestringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiLinestringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiLinestringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLinestringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLinestringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiLinestring(s)
	}
}

func (s *MultiLinestringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiLinestring(s)
	}
}

func (p *CQLParser) MultiLinestring() (localctx IMultiLinestringContext) {
	localctx = NewMultiLinestringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CQLParserRULE_multiLinestring)
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
		p.SetState(260)
		p.Match(CQLParserMULTILINESTRING)
	}
	{
		p.SetState(261)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(262)
		p.CoordList()
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(263)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(264)
			p.CoordList()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(270)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonContext() *MultiPolygonContext {
	var p = new(MultiPolygonContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_multiPolygon
	return p
}

func (*MultiPolygonContext) IsMultiPolygonContext() {}

func NewMultiPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonContext {
	var p = new(MultiPolygonContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_multiPolygon

	return p
}

func (s *MultiPolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(CQLParserMULTIPOLYGON, 0)
}

func (s *MultiPolygonContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
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
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *MultiPolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *MultiPolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *MultiPolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterMultiPolygon(s)
	}
}

func (s *MultiPolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitMultiPolygon(s)
	}
}

func (p *CQLParser) MultiPolygon() (localctx IMultiPolygonContext) {
	localctx = NewMultiPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CQLParserRULE_multiPolygon)
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
		p.SetState(272)
		p.Match(CQLParserMULTIPOLYGON)
	}
	{
		p.SetState(273)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(274)
		p.PolygonDef()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(275)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(276)
			p.PolygonDef()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyGeometryCollectionContext() *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_geometryCollection
	return p
}

func (*GeometryCollectionContext) IsGeometryCollectionContext() {}

func NewGeometryCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_geometryCollection

	return p
}

func (s *GeometryCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryCollectionContext) GEOMETRYCOLLECTION() antlr.TerminalNode {
	return s.GetToken(CQLParserGEOMETRYCOLLECTION, 0)
}

func (s *GeometryCollectionContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
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
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *GeometryCollectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *GeometryCollectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *GeometryCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterGeometryCollection(s)
	}
}

func (s *GeometryCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitGeometryCollection(s)
	}
}

func (p *CQLParser) GeometryCollection() (localctx IGeometryCollectionContext) {
	localctx = NewGeometryCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CQLParserRULE_geometryCollection)
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
		p.SetState(284)
		p.Match(CQLParserGEOMETRYCOLLECTION)
	}
	{
		p.SetState(285)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(286)
		p.GeomLiteral()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(287)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(288)
			p.GeomLiteral()
		}

		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(294)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyEnvelopeContext() *EnvelopeContext {
	var p = new(EnvelopeContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_envelope
	return p
}

func (*EnvelopeContext) IsEnvelopeContext() {}

func NewEnvelopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvelopeContext {
	var p = new(EnvelopeContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_envelope

	return p
}

func (s *EnvelopeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvelopeContext) ENVELOPE() antlr.TerminalNode {
	return s.GetToken(CQLParserENVELOPE, 0)
}

func (s *EnvelopeContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *EnvelopeContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *EnvelopeContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *EnvelopeContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *EnvelopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvelopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvelopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterEnvelope(s)
	}
}

func (s *EnvelopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitEnvelope(s)
	}
}

func (p *CQLParser) Envelope() (localctx IEnvelopeContext) {
	localctx = NewEnvelopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CQLParserRULE_envelope)

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
		p.Match(CQLParserENVELOPE)
	}
	{
		p.SetState(297)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(298)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(299)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(300)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(301)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(302)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(303)
		p.Match(CQLParserCOMMA)
	}
	{
		p.SetState(304)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(305)
		p.Match(CQLParserRIGHTPAREN)
	}

	return localctx
}

// ICoordListContext is an interface to support dynamic dispatch.
type ICoordListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordListContext differentiates from other interfaces.
	IsCoordListContext()
}

type CoordListContext struct {
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordListContext() *CoordListContext {
	var p = new(CoordListContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_coordList
	return p
}

func (*CoordListContext) IsCoordListContext() {}

func NewCoordListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordListContext {
	var p = new(CoordListContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordList

	return p
}

func (s *CoordListContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserLEFTPAREN, 0)
}

func (s *CoordListContext) AllCoordinate() []ICoordinateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICoordinateContext)(nil)).Elem())
	var tst = make([]ICoordinateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICoordinateContext)
		}
	}

	return tst
}

func (s *CoordListContext) Coordinate(i int) ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *CoordListContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLParserRIGHTPAREN, 0)
}

func (s *CoordListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLParserCOMMA)
}

func (s *CoordListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserCOMMA, i)
}

func (s *CoordListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordList(s)
	}
}

func (s *CoordListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordList(s)
	}
}

func (p *CQLParser) CoordList() (localctx ICoordListContext) {
	localctx = NewCoordListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CQLParserRULE_coordList)
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
		p.SetState(307)
		p.Match(CQLParserLEFTPAREN)
	}
	{
		p.SetState(308)
		p.Coordinate()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLParserCOMMA {
		{
			p.SetState(309)
			p.Match(CQLParserCOMMA)
		}
		{
			p.SetState(310)
			p.Coordinate()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(316)
		p.Match(CQLParserRIGHTPAREN)
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
	*CqlContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.CqlContext = NewCqlContext(nil, -1)
	p.RuleIndex = CQLParserRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.CqlContext = NewCqlContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLParserNumericLiteral)
}

func (s *CoordinateContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLParserNumericLiteral, i)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.EnterCoordinate(s)
	}
}

func (s *CoordinateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLParserListener); ok {
		listenerT.ExitCoordinate(s)
	}
}

func (p *CQLParser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CQLParserRULE_coordinate)

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
		p.SetState(318)
		p.Match(CQLParserNumericLiteral)
	}
	{
		p.SetState(319)
		p.Match(CQLParserNumericLiteral)
	}

	return localctx
}

func (p *CQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ScalarExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ScalarExpressionContext)
		}
		return p.ScalarExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CQLParser) ScalarExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
