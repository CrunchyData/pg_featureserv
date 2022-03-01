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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 93, 319,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 7, 3, 80, 10, 3, 12, 3, 14, 3, 83, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 7, 4, 91, 10, 4, 12, 4, 14, 4, 94, 11, 4, 3, 5, 5, 5, 97, 10,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 106, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 115, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 5, 9, 123, 10, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10,
	130, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5,
	11, 140, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 148,
	10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 5, 19, 176, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 186, 10, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 7, 23, 200, 10, 23, 12, 23, 14, 23, 203, 11, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 7, 25, 214, 10, 25, 12, 25,
	14, 25, 217, 11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7,
	26, 226, 10, 26, 12, 26, 14, 26, 229, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 7, 27, 238, 10, 27, 12, 27, 14, 27, 241, 11, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28, 250, 10, 28, 12, 28,
	14, 28, 253, 11, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7,
	29, 262, 10, 29, 12, 29, 14, 29, 265, 11, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 5, 33, 289, 10, 33,
	3, 34, 3, 34, 3, 35, 3, 35, 5, 35, 295, 10, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 7, 35, 302, 10, 35, 12, 35, 14, 35, 305, 11, 35, 3, 35, 3, 35,
	3, 35, 7, 35, 310, 10, 35, 12, 35, 14, 35, 313, 11, 35, 5, 35, 315, 10,
	35, 3, 35, 3, 35, 3, 35, 2, 4, 4, 6, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
	56, 58, 60, 62, 64, 66, 68, 2, 4, 3, 2, 14, 15, 4, 2, 3, 3, 25, 25, 2,
	319, 2, 70, 3, 2, 2, 2, 4, 73, 3, 2, 2, 2, 6, 84, 3, 2, 2, 2, 8, 96, 3,
	2, 2, 2, 10, 105, 3, 2, 2, 2, 12, 114, 3, 2, 2, 2, 14, 116, 3, 2, 2, 2,
	16, 120, 3, 2, 2, 2, 18, 127, 3, 2, 2, 2, 20, 136, 3, 2, 2, 2, 22, 147,
	3, 2, 2, 2, 24, 149, 3, 2, 2, 2, 26, 151, 3, 2, 2, 2, 28, 153, 3, 2, 2,
	2, 30, 155, 3, 2, 2, 2, 32, 157, 3, 2, 2, 2, 34, 164, 3, 2, 2, 2, 36, 175,
	3, 2, 2, 2, 38, 185, 3, 2, 2, 2, 40, 187, 3, 2, 2, 2, 42, 192, 3, 2, 2,
	2, 44, 195, 3, 2, 2, 2, 46, 206, 3, 2, 2, 2, 48, 209, 3, 2, 2, 2, 50, 220,
	3, 2, 2, 2, 52, 232, 3, 2, 2, 2, 54, 244, 3, 2, 2, 2, 56, 256, 3, 2, 2,
	2, 58, 268, 3, 2, 2, 2, 60, 279, 3, 2, 2, 2, 62, 282, 3, 2, 2, 2, 64, 288,
	3, 2, 2, 2, 66, 290, 3, 2, 2, 2, 68, 292, 3, 2, 2, 2, 70, 71, 5, 4, 3,
	2, 71, 72, 7, 2, 2, 3, 72, 3, 3, 2, 2, 2, 73, 74, 8, 3, 1, 2, 74, 75, 5,
	6, 4, 2, 75, 81, 3, 2, 2, 2, 76, 77, 12, 3, 2, 2, 77, 78, 7, 12, 2, 2,
	78, 80, 5, 6, 4, 2, 79, 76, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3,
	2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 5, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84,
	85, 8, 4, 1, 2, 85, 86, 5, 8, 5, 2, 86, 92, 3, 2, 2, 2, 87, 88, 12, 3,
	2, 2, 88, 89, 7, 11, 2, 2, 89, 91, 5, 8, 5, 2, 90, 87, 3, 2, 2, 2, 91,
	94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 7, 3, 2, 2,
	2, 94, 92, 3, 2, 2, 2, 95, 97, 7, 13, 2, 2, 96, 95, 3, 2, 2, 2, 96, 97,
	3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 5, 10, 6, 2, 99, 9, 3, 2, 2, 2,
	100, 106, 5, 12, 7, 2, 101, 102, 7, 52, 2, 2, 102, 103, 5, 4, 3, 2, 103,
	104, 7, 53, 2, 2, 104, 106, 3, 2, 2, 2, 105, 100, 3, 2, 2, 2, 105, 101,
	3, 2, 2, 2, 106, 11, 3, 2, 2, 2, 107, 115, 5, 14, 8, 2, 108, 115, 5, 16,
	9, 2, 109, 115, 5, 18, 10, 2, 110, 115, 5, 20, 11, 2, 111, 115, 5, 68,
	35, 2, 112, 115, 5, 32, 17, 2, 113, 115, 5, 34, 18, 2, 114, 107, 3, 2,
	2, 2, 114, 108, 3, 2, 2, 2, 114, 109, 3, 2, 2, 2, 114, 110, 3, 2, 2, 2,
	114, 111, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 113, 3, 2, 2, 2, 115,
	13, 3, 2, 2, 2, 116, 117, 5, 22, 12, 2, 117, 118, 7, 3, 2, 2, 118, 119,
	5, 22, 12, 2, 119, 15, 3, 2, 2, 2, 120, 122, 5, 24, 13, 2, 121, 123, 7,
	13, 2, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2,
	2, 124, 125, 9, 2, 2, 2, 125, 126, 5, 26, 14, 2, 126, 17, 3, 2, 2, 2, 127,
	129, 5, 24, 13, 2, 128, 130, 7, 13, 2, 2, 129, 128, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 16, 2, 2, 132, 133, 5, 22,
	12, 2, 133, 134, 7, 11, 2, 2, 134, 135, 5, 22, 12, 2, 135, 19, 3, 2, 2,
	2, 136, 137, 5, 24, 13, 2, 137, 139, 7, 17, 2, 2, 138, 140, 7, 13, 2, 2,
	139, 138, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141,
	142, 7, 18, 2, 2, 142, 21, 3, 2, 2, 2, 143, 148, 5, 24, 13, 2, 144, 148,
	5, 26, 14, 2, 145, 148, 5, 28, 15, 2, 146, 148, 5, 30, 16, 2, 147, 143,
	3, 2, 2, 2, 147, 144, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 146, 3, 2,
	2, 2, 148, 23, 3, 2, 2, 2, 149, 150, 7, 40, 2, 2, 150, 25, 3, 2, 2, 2,
	151, 152, 7, 92, 2, 2, 152, 27, 3, 2, 2, 2, 153, 154, 7, 39, 2, 2, 154,
	29, 3, 2, 2, 2, 155, 156, 7, 10, 2, 2, 156, 31, 3, 2, 2, 2, 157, 158, 7,
	23, 2, 2, 158, 159, 7, 52, 2, 2, 159, 160, 5, 36, 19, 2, 160, 161, 7, 58,
	2, 2, 161, 162, 5, 36, 19, 2, 162, 163, 7, 53, 2, 2, 163, 33, 3, 2, 2,
	2, 164, 165, 7, 24, 2, 2, 165, 166, 7, 52, 2, 2, 166, 167, 5, 36, 19, 2,
	167, 168, 7, 58, 2, 2, 168, 169, 5, 36, 19, 2, 169, 170, 7, 58, 2, 2, 170,
	171, 7, 39, 2, 2, 171, 172, 7, 53, 2, 2, 172, 35, 3, 2, 2, 2, 173, 176,
	5, 24, 13, 2, 174, 176, 5, 38, 20, 2, 175, 173, 3, 2, 2, 2, 175, 174, 3,
	2, 2, 2, 176, 37, 3, 2, 2, 2, 177, 186, 5, 40, 21, 2, 178, 186, 5, 42,
	22, 2, 179, 186, 5, 46, 24, 2, 180, 186, 5, 50, 26, 2, 181, 186, 5, 52,
	27, 2, 182, 186, 5, 54, 28, 2, 183, 186, 5, 56, 29, 2, 184, 186, 5, 58,
	30, 2, 185, 177, 3, 2, 2, 2, 185, 178, 3, 2, 2, 2, 185, 179, 3, 2, 2, 2,
	185, 180, 3, 2, 2, 2, 185, 181, 3, 2, 2, 2, 185, 182, 3, 2, 2, 2, 185,
	183, 3, 2, 2, 2, 185, 184, 3, 2, 2, 2, 186, 39, 3, 2, 2, 2, 187, 188, 7,
	31, 2, 2, 188, 189, 7, 52, 2, 2, 189, 190, 5, 60, 31, 2, 190, 191, 7, 53,
	2, 2, 191, 41, 3, 2, 2, 2, 192, 193, 7, 32, 2, 2, 193, 194, 5, 44, 23,
	2, 194, 43, 3, 2, 2, 2, 195, 196, 7, 52, 2, 2, 196, 201, 5, 60, 31, 2,
	197, 198, 7, 58, 2, 2, 198, 200, 5, 60, 31, 2, 199, 197, 3, 2, 2, 2, 200,
	203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204,
	3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 204, 205, 7, 53, 2, 2, 205, 45, 3, 2,
	2, 2, 206, 207, 7, 33, 2, 2, 207, 208, 5, 48, 25, 2, 208, 47, 3, 2, 2,
	2, 209, 210, 7, 52, 2, 2, 210, 215, 5, 44, 23, 2, 211, 212, 7, 58, 2, 2,
	212, 214, 5, 44, 23, 2, 213, 211, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215,
	213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218, 3, 2, 2, 2, 217, 215,
	3, 2, 2, 2, 218, 219, 7, 53, 2, 2, 219, 49, 3, 2, 2, 2, 220, 221, 7, 34,
	2, 2, 221, 222, 7, 52, 2, 2, 222, 227, 5, 60, 31, 2, 223, 224, 7, 58, 2,
	2, 224, 226, 5, 60, 31, 2, 225, 223, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2,
	227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229,
	227, 3, 2, 2, 2, 230, 231, 7, 53, 2, 2, 231, 51, 3, 2, 2, 2, 232, 233,
	7, 35, 2, 2, 233, 234, 7, 52, 2, 2, 234, 239, 5, 44, 23, 2, 235, 236, 7,
	58, 2, 2, 236, 238, 5, 44, 23, 2, 237, 235, 3, 2, 2, 2, 238, 241, 3, 2,
	2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 242, 3, 2, 2, 2,
	241, 239, 3, 2, 2, 2, 242, 243, 7, 53, 2, 2, 243, 53, 3, 2, 2, 2, 244,
	245, 7, 36, 2, 2, 245, 246, 7, 52, 2, 2, 246, 251, 5, 48, 25, 2, 247, 248,
	7, 58, 2, 2, 248, 250, 5, 48, 25, 2, 249, 247, 3, 2, 2, 2, 250, 253, 3,
	2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 254, 3, 2, 2,
	2, 253, 251, 3, 2, 2, 2, 254, 255, 7, 53, 2, 2, 255, 55, 3, 2, 2, 2, 256,
	257, 7, 37, 2, 2, 257, 258, 7, 52, 2, 2, 258, 263, 5, 38, 20, 2, 259, 260,
	7, 58, 2, 2, 260, 262, 5, 38, 20, 2, 261, 259, 3, 2, 2, 2, 262, 265, 3,
	2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 266, 3, 2, 2,
	2, 265, 263, 3, 2, 2, 2, 266, 267, 7, 53, 2, 2, 267, 57, 3, 2, 2, 2, 268,
	269, 7, 38, 2, 2, 269, 270, 7, 52, 2, 2, 270, 271, 7, 39, 2, 2, 271, 272,
	7, 58, 2, 2, 272, 273, 7, 39, 2, 2, 273, 274, 7, 58, 2, 2, 274, 275, 7,
	39, 2, 2, 275, 276, 7, 58, 2, 2, 276, 277, 7, 39, 2, 2, 277, 278, 7, 53,
	2, 2, 278, 59, 3, 2, 2, 2, 279, 280, 7, 39, 2, 2, 280, 281, 7, 39, 2, 2,
	281, 61, 3, 2, 2, 2, 282, 283, 5, 64, 33, 2, 283, 284, 9, 3, 2, 2, 284,
	285, 5, 64, 33, 2, 285, 63, 3, 2, 2, 2, 286, 289, 5, 24, 13, 2, 287, 289,
	5, 66, 34, 2, 288, 286, 3, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289, 65, 3, 2,
	2, 2, 290, 291, 7, 77, 2, 2, 291, 67, 3, 2, 2, 2, 292, 294, 5, 24, 13,
	2, 293, 295, 7, 13, 2, 2, 294, 293, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295,
	296, 3, 2, 2, 2, 296, 297, 7, 30, 2, 2, 297, 314, 7, 52, 2, 2, 298, 303,
	5, 26, 14, 2, 299, 300, 7, 58, 2, 2, 300, 302, 5, 26, 14, 2, 301, 299,
	3, 2, 2, 2, 302, 305, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2,
	2, 2, 304, 315, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 306, 311, 5, 28, 15,
	2, 307, 308, 7, 58, 2, 2, 308, 310, 5, 28, 15, 2, 309, 307, 3, 2, 2, 2,
	310, 313, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312,
	315, 3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 314, 298, 3, 2, 2, 2, 314, 306,
	3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 317, 7, 53, 2, 2, 317, 69, 3, 2,
	2, 2, 24, 81, 92, 96, 105, 114, 122, 129, 139, 147, 175, 185, 201, 215,
	227, 239, 251, 263, 288, 294, 303, 311, 314,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'<'", "'='", "'>'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "'#'", "'$'", "'_'", "'\"'", "'%'",
	"'&'", "", "'('", "')'", "'['", "']'", "'*'", "'+'", "','", "'-'", "'.'",
	"'/'", "':'", "';'", "'?'", "'|'", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"''''",
}
var symbolicNames = []string{
	"", "ComparisonOperator", "LT", "EQ", "GT", "NEQ", "GTEQ", "LTEQ", "BooleanLiteral",
	"AND", "OR", "NOT", "LIKE", "ILIKE", "BETWEEN", "IS", "NULL", "WILDCARD",
	"SINGLECHAR", "ESCAPECHAR", "NOCASE", "SpatialOperator", "DistanceOperator",
	"TemporalOperator", "ArrayOperator", "EXISTS", "EXIST", "DOES", "IN", "POINT",
	"LINESTRING", "POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON",
	"GEOMETRYCOLLECTION", "ENVELOPE", "NumericLiteral", "Identifier", "IdentifierStart",
	"IdentifierPart", "ALPHA", "DIGIT", "OCTOTHORP", "DOLLAR", "UNDERSCORE",
	"DOUBLEQUOTE", "PERCENT", "AMPERSAND", "QUOTE", "LEFTPAREN", "RIGHTPAREN",
	"LEFTSQUAREBRACKET", "RIGHTSQUAREBRACKET", "ASTERISK", "PLUS", "COMMA",
	"MINUS", "PERIOD", "SOLIDUS", "COLON", "SEMICOLON", "QUESTIONMARK", "VERTICALBAR",
	"BIT", "HEXIT", "UnsignedNumericLiteral", "SignedNumericLiteral", "ExactNumericLiteral",
	"ApproximateNumericLiteral", "Mantissa", "Exponent", "SignedInteger", "UnsignedInteger",
	"Sign", "TemporalLiteral", "Instant", "Interval", "InstantInInterval",
	"FullDate", "DateYear", "DateMonth", "DateDay", "UtcTime", "TimeZoneOffset",
	"TimeHour", "TimeMinute", "TimeSecond", "NOW", "WS", "CharacterStringLiteral",
	"QuotedQuote",
}

var ruleNames = []string{
	"cqlFilter", "booleanValueExpression", "booleanTerm", "booleanFactor",
	"booleanPrimary", "predicate", "binaryComparisonPredicate", "likePredicate",
	"betweenPredicate", "isNullPredicate", "scalarExpression", "propertyName",
	"characterLiteral", "numericLiteral", "booleanLiteral", "spatialPredicate",
	"distancePredicate", "geomExpression", "geomLiteral", "point", "linestring",
	"coordList", "polygon", "polygonDef", "multiPoint", "multiLinestring",
	"multiPolygon", "geometryCollection", "envelope", "coordinate", "temporalPredicate",
	"temporalExpression", "temporalLiteral", "inPredicate",
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
	CQLDistanceOperator          = 22
	CQLTemporalOperator          = 23
	CQLArrayOperator             = 24
	CQLEXISTS                    = 25
	CQLEXIST                     = 26
	CQLDOES                      = 27
	CQLIN                        = 28
	CQLPOINT                     = 29
	CQLLINESTRING                = 30
	CQLPOLYGON                   = 31
	CQLMULTIPOINT                = 32
	CQLMULTILINESTRING           = 33
	CQLMULTIPOLYGON              = 34
	CQLGEOMETRYCOLLECTION        = 35
	CQLENVELOPE                  = 36
	CQLNumericLiteral            = 37
	CQLIdentifier                = 38
	CQLIdentifierStart           = 39
	CQLIdentifierPart            = 40
	CQLALPHA                     = 41
	CQLDIGIT                     = 42
	CQLOCTOTHORP                 = 43
	CQLDOLLAR                    = 44
	CQLUNDERSCORE                = 45
	CQLDOUBLEQUOTE               = 46
	CQLPERCENT                   = 47
	CQLAMPERSAND                 = 48
	CQLQUOTE                     = 49
	CQLLEFTPAREN                 = 50
	CQLRIGHTPAREN                = 51
	CQLLEFTSQUAREBRACKET         = 52
	CQLRIGHTSQUAREBRACKET        = 53
	CQLASTERISK                  = 54
	CQLPLUS                      = 55
	CQLCOMMA                     = 56
	CQLMINUS                     = 57
	CQLPERIOD                    = 58
	CQLSOLIDUS                   = 59
	CQLCOLON                     = 60
	CQLSEMICOLON                 = 61
	CQLQUESTIONMARK              = 62
	CQLVERTICALBAR               = 63
	CQLBIT                       = 64
	CQLHEXIT                     = 65
	CQLUnsignedNumericLiteral    = 66
	CQLSignedNumericLiteral      = 67
	CQLExactNumericLiteral       = 68
	CQLApproximateNumericLiteral = 69
	CQLMantissa                  = 70
	CQLExponent                  = 71
	CQLSignedInteger             = 72
	CQLUnsignedInteger           = 73
	CQLSign                      = 74
	CQLTemporalLiteral           = 75
	CQLInstant                   = 76
	CQLInterval                  = 77
	CQLInstantInInterval         = 78
	CQLFullDate                  = 79
	CQLDateYear                  = 80
	CQLDateMonth                 = 81
	CQLDateDay                   = 82
	CQLUtcTime                   = 83
	CQLTimeZoneOffset            = 84
	CQLTimeHour                  = 85
	CQLTimeMinute                = 86
	CQLTimeSecond                = 87
	CQLNOW                       = 88
	CQLWS                        = 89
	CQLCharacterStringLiteral    = 90
	CQLQuotedQuote               = 91
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
	CQLRULE_distancePredicate         = 16
	CQLRULE_geomExpression            = 17
	CQLRULE_geomLiteral               = 18
	CQLRULE_point                     = 19
	CQLRULE_linestring                = 20
	CQLRULE_coordList                 = 21
	CQLRULE_polygon                   = 22
	CQLRULE_polygonDef                = 23
	CQLRULE_multiPoint                = 24
	CQLRULE_multiLinestring           = 25
	CQLRULE_multiPolygon              = 26
	CQLRULE_geometryCollection        = 27
	CQLRULE_envelope                  = 28
	CQLRULE_coordinate                = 29
	CQLRULE_temporalPredicate         = 30
	CQLRULE_temporalExpression        = 31
	CQLRULE_temporalLiteral           = 32
	CQLRULE_inPredicate               = 33
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
		p.SetState(68)
		p.booleanValueExpression(0)
	}
	{
		p.SetState(69)
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
		p.SetState(72)
		p.booleanTerm(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
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
			p.SetState(74)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(75)
				p.Match(CQLOR)
			}
			{
				p.SetState(76)
				p.booleanTerm(0)
			}

		}
		p.SetState(81)
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
		p.SetState(83)
		p.BooleanFactor()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
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
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(86)
				p.Match(CQLAND)
			}
			{
				p.SetState(87)
				p.BooleanFactor()
			}

		}
		p.SetState(92)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(93)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(96)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLBooleanLiteral, CQLSpatialOperator, CQLDistanceOperator, CQLNumericLiteral, CQLIdentifier, CQLCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Predicate()
		}

	case CQLLEFTPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(CQLLEFTPAREN)
		}
		{
			p.SetState(100)
			p.booleanValueExpression(0)
		}
		{
			p.SetState(101)
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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.BinaryComparisonPredicate()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.LikePredicate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.BetweenPredicate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.IsNullPredicate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)
			p.InPredicate()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(110)
			p.SpatialPredicate()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(111)
			p.DistancePredicate()
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
		p.SetState(114)
		p.ScalarExpression()
	}
	{
		p.SetState(115)
		p.Match(CQLComparisonOperator)
	}
	{
		p.SetState(116)
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
		p.SetState(118)
		p.PropertyName()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(119)
			p.Match(CQLNOT)
		}

	}
	p.SetState(122)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLLIKE || _la == CQLILIKE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(123)
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
		p.SetState(125)
		p.PropertyName()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(126)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(129)
		p.Match(CQLBETWEEN)
	}
	{
		p.SetState(130)
		p.ScalarExpression()
	}
	{
		p.SetState(131)
		p.Match(CQLAND)
	}
	{
		p.SetState(132)
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
		p.SetState(134)
		p.PropertyName()
	}
	{
		p.SetState(135)
		p.Match(CQLIS)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(136)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(139)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.PropertyName()
		}

	case CQLCharacterStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.CharacterLiteral()
		}

	case CQLNumericLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(143)
			p.NumericLiteral()
		}

	case CQLBooleanLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)
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
		p.SetState(147)
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
		p.SetState(149)
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
		p.SetState(151)
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
		p.SetState(153)
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
		p.SetState(155)
		p.Match(CQLSpatialOperator)
	}
	{
		p.SetState(156)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(157)
		p.GeomExpression()
	}
	{
		p.SetState(158)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(159)
		p.GeomExpression()
	}
	{
		p.SetState(160)
		p.Match(CQLRIGHTPAREN)
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistancePredicateContext() *DistancePredicateContext {
	var p = new(DistancePredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_distancePredicate
	return p
}

func (*DistancePredicateContext) IsDistancePredicateContext() {}

func NewDistancePredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistancePredicateContext {
	var p = new(DistancePredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_distancePredicate

	return p
}

func (s *DistancePredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DistancePredicateContext) DistanceOperator() antlr.TerminalNode {
	return s.GetToken(CQLDistanceOperator, 0)
}

func (s *DistancePredicateContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
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
	return s.GetTokens(CQLCOMMA)
}

func (s *DistancePredicateContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *DistancePredicateContext) NumericLiteral() antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, 0)
}

func (s *DistancePredicateContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *DistancePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistancePredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistancePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterDistancePredicate(s)
	}
}

func (s *DistancePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitDistancePredicate(s)
	}
}

func (p *CQL) DistancePredicate() (localctx IDistancePredicateContext) {
	localctx = NewDistancePredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CQLRULE_distancePredicate)

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
		p.SetState(162)
		p.Match(CQLDistanceOperator)
	}
	{
		p.SetState(163)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(164)
		p.GeomExpression()
	}
	{
		p.SetState(165)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(166)
		p.GeomExpression()
	}
	{
		p.SetState(167)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(168)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(169)
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
	p.EnterRule(localctx, 34, CQLRULE_geomExpression)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.PropertyName()
		}

	case CQLPOINT, CQLLINESTRING, CQLPOLYGON, CQLMULTIPOINT, CQLMULTILINESTRING, CQLMULTIPOLYGON, CQLGEOMETRYCOLLECTION, CQLENVELOPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
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
	p.EnterRule(localctx, 36, CQLRULE_geomLiteral)

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

	p.SetState(183)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLPOINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Point()
		}

	case CQLLINESTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Linestring()
		}

	case CQLPOLYGON:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Polygon()
		}

	case CQLMULTIPOINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(178)
			p.MultiPoint()
		}

	case CQLMULTILINESTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(179)
			p.MultiLinestring()
		}

	case CQLMULTIPOLYGON:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.MultiPolygon()
		}

	case CQLGEOMETRYCOLLECTION:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(181)
			p.GeometryCollection()
		}

	case CQLENVELOPE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(182)
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
	p.EnterRule(localctx, 38, CQLRULE_point)

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
		p.SetState(185)
		p.Match(CQLPOINT)
	}
	{
		p.SetState(186)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(187)
		p.Coordinate()
	}
	{
		p.SetState(188)
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
	p.EnterRule(localctx, 40, CQLRULE_linestring)

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
		p.Match(CQLLINESTRING)
	}
	{
		p.SetState(191)
		p.CoordList()
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordListContext() *CoordListContext {
	var p = new(CoordListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CQLRULE_coordList
	return p
}

func (*CoordListContext) IsCoordListContext() {}

func NewCoordListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordListContext {
	var p = new(CoordListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CQLRULE_coordList

	return p
}

func (s *CoordListContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordListContext) LEFTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLLEFTPAREN, 0)
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
	return s.GetToken(CQLRIGHTPAREN, 0)
}

func (s *CoordListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *CoordListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *CoordListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.EnterCoordList(s)
	}
}

func (s *CoordListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CQLListener); ok {
		listenerT.ExitCoordList(s)
	}
}

func (p *CQL) CoordList() (localctx ICoordListContext) {
	localctx = NewCoordListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CQLRULE_coordList)
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
		p.SetState(193)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(194)
		p.Coordinate()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(195)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(196)
			p.Coordinate()
		}

		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
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
	p.EnterRule(localctx, 44, CQLRULE_polygon)

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
		p.SetState(204)
		p.Match(CQLPOLYGON)
	}
	{
		p.SetState(205)
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
	p.EnterRule(localctx, 46, CQLRULE_polygonDef)
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
		p.SetState(207)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(208)
		p.CoordList()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(209)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(210)
			p.CoordList()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
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
	p.EnterRule(localctx, 48, CQLRULE_multiPoint)
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
		p.SetState(218)
		p.Match(CQLMULTIPOINT)
	}
	{
		p.SetState(219)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(220)
		p.Coordinate()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(221)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(222)
			p.Coordinate()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(228)
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
	p.EnterRule(localctx, 50, CQLRULE_multiLinestring)
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
		p.SetState(230)
		p.Match(CQLMULTILINESTRING)
	}
	{
		p.SetState(231)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(232)
		p.CoordList()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(233)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(234)
			p.CoordList()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(240)
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
	p.EnterRule(localctx, 52, CQLRULE_multiPolygon)
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
		p.SetState(242)
		p.Match(CQLMULTIPOLYGON)
	}
	{
		p.SetState(243)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(244)
		p.PolygonDef()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(245)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(246)
			p.PolygonDef()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
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
	p.EnterRule(localctx, 54, CQLRULE_geometryCollection)
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
		p.SetState(254)
		p.Match(CQLGEOMETRYCOLLECTION)
	}
	{
		p.SetState(255)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(256)
		p.GeomLiteral()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CQLCOMMA {
		{
			p.SetState(257)
			p.Match(CQLCOMMA)
		}
		{
			p.SetState(258)
			p.GeomLiteral()
		}

		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(264)
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

func (s *EnvelopeContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLNumericLiteral)
}

func (s *EnvelopeContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, i)
}

func (s *EnvelopeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CQLCOMMA)
}

func (s *EnvelopeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CQLCOMMA, i)
}

func (s *EnvelopeContext) RIGHTPAREN() antlr.TerminalNode {
	return s.GetToken(CQLRIGHTPAREN, 0)
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
	p.EnterRule(localctx, 56, CQLRULE_envelope)

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
		p.SetState(266)
		p.Match(CQLENVELOPE)
	}
	{
		p.SetState(267)
		p.Match(CQLLEFTPAREN)
	}
	{
		p.SetState(268)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(269)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(270)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(271)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(272)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(273)
		p.Match(CQLCOMMA)
	}
	{
		p.SetState(274)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(275)
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

func (s *CoordinateContext) AllNumericLiteral() []antlr.TerminalNode {
	return s.GetTokens(CQLNumericLiteral)
}

func (s *CoordinateContext) NumericLiteral(i int) antlr.TerminalNode {
	return s.GetToken(CQLNumericLiteral, i)
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
	p.EnterRule(localctx, 58, CQLRULE_coordinate)

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
		p.SetState(277)
		p.Match(CQLNumericLiteral)
	}
	{
		p.SetState(278)
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
	p.EnterRule(localctx, 60, CQLRULE_temporalPredicate)
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
		p.SetState(280)
		p.TemporalExpression()
	}
	p.SetState(281)
	_la = p.GetTokenStream().LA(1)

	if !(_la == CQLComparisonOperator || _la == CQLTemporalOperator) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(282)
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
	p.EnterRule(localctx, 62, CQLRULE_temporalExpression)

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

	p.SetState(286)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.PropertyName()
		}

	case CQLTemporalLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
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
	p.EnterRule(localctx, 64, CQLRULE_temporalLiteral)

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
		p.SetState(288)
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
	p.EnterRule(localctx, 66, CQLRULE_inPredicate)
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
		p.SetState(290)
		p.PropertyName()
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CQLNOT {
		{
			p.SetState(291)
			p.Match(CQLNOT)
		}

	}
	{
		p.SetState(294)
		p.Match(CQLIN)
	}
	{
		p.SetState(295)
		p.Match(CQLLEFTPAREN)
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CQLCharacterStringLiteral:
		{
			p.SetState(296)
			p.CharacterLiteral()
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLCOMMA {
			{
				p.SetState(297)
				p.Match(CQLCOMMA)
			}
			{
				p.SetState(298)
				p.CharacterLiteral()
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case CQLNumericLiteral:
		{
			p.SetState(304)
			p.NumericLiteral()
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == CQLCOMMA {
			{
				p.SetState(305)
				p.Match(CQLCOMMA)
			}
			{
				p.SetState(306)
				p.NumericLiteral()
			}

			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(314)
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
