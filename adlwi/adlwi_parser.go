// Generated from AdlWi.g4 by ANTLR 4.7.

package adlw1 // AdlWi
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/wxio/goantlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 301,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 4, 7, 4, 45, 10, 4, 12,
	4, 14, 4, 48, 11, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 3,
	4, 5, 4, 57, 10, 4, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6,
	66, 10, 6, 12, 6, 14, 6, 69, 11, 6, 3, 6, 5, 6, 72, 10, 6, 3, 6, 7, 6,
	75, 10, 6, 12, 6, 14, 6, 78, 11, 6, 3, 6, 5, 6, 81, 10, 6, 3, 6, 3, 6,
	3, 6, 7, 6, 86, 10, 6, 12, 6, 14, 6, 89, 11, 6, 3, 6, 5, 6, 92, 10, 6,
	3, 6, 7, 6, 95, 10, 6, 12, 6, 14, 6, 98, 11, 6, 3, 6, 5, 6, 101, 10, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 106, 10, 6, 12, 6, 14, 6, 109, 11, 6, 3, 6, 5,
	6, 112, 10, 6, 3, 6, 5, 6, 115, 10, 6, 3, 6, 7, 6, 118, 10, 6, 12, 6, 14,
	6, 121, 11, 6, 3, 6, 5, 6, 124, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 129, 10,
	6, 12, 6, 14, 6, 132, 11, 6, 3, 6, 5, 6, 135, 10, 6, 3, 6, 5, 6, 138, 10,
	6, 3, 6, 7, 6, 141, 10, 6, 12, 6, 14, 6, 144, 11, 6, 3, 6, 5, 6, 147, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 166, 10, 6, 3, 6, 3, 6, 7, 6, 170,
	10, 6, 12, 6, 14, 6, 173, 11, 6, 3, 6, 3, 6, 7, 6, 177, 10, 6, 12, 6, 14,
	6, 180, 11, 6, 3, 6, 3, 6, 3, 6, 5, 6, 185, 10, 6, 3, 6, 3, 6, 7, 6, 189,
	10, 6, 12, 6, 14, 6, 192, 11, 6, 3, 6, 3, 6, 5, 6, 196, 10, 6, 3, 6, 7,
	6, 199, 10, 6, 12, 6, 14, 6, 202, 11, 6, 3, 6, 5, 6, 205, 10, 6, 5, 6,
	207, 10, 6, 3, 7, 3, 7, 3, 7, 7, 7, 212, 10, 7, 12, 7, 14, 7, 215, 11,
	7, 3, 7, 5, 7, 218, 10, 7, 3, 7, 5, 7, 221, 10, 7, 3, 7, 5, 7, 224, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 231, 10, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 238, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 245, 10,
	8, 5, 8, 247, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 253, 10, 9, 13, 9, 14,
	9, 254, 3, 9, 3, 9, 5, 9, 259, 10, 9, 3, 10, 3, 10, 3, 10, 6, 10, 264,
	10, 10, 13, 10, 14, 10, 265, 3, 10, 3, 10, 5, 10, 270, 10, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 280, 10, 11, 13, 11,
	14, 11, 281, 3, 11, 3, 11, 5, 11, 286, 10, 11, 3, 11, 3, 11, 3, 11, 6,
	11, 291, 10, 11, 13, 11, 14, 11, 292, 3, 11, 3, 11, 5, 11, 297, 10, 11,
	5, 11, 299, 10, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	2, 2, 2, 352, 2, 22, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2,
	8, 60, 3, 2, 2, 2, 10, 206, 3, 2, 2, 2, 12, 208, 3, 2, 2, 2, 14, 246, 3,
	2, 2, 2, 16, 258, 3, 2, 2, 2, 18, 260, 3, 2, 2, 2, 20, 298, 3, 2, 2, 2,
	22, 23, 7, 26, 2, 2, 23, 24, 7, 30, 2, 2, 24, 25, 7, 26, 2, 2, 25, 26,
	5, 6, 4, 2, 26, 27, 7, 27, 2, 2, 27, 28, 7, 27, 2, 2, 28, 29, 7, 2, 2,
	3, 29, 3, 3, 2, 2, 2, 30, 31, 7, 26, 2, 2, 31, 32, 5, 20, 11, 2, 32, 33,
	7, 27, 2, 2, 33, 34, 7, 2, 2, 3, 34, 5, 3, 2, 2, 2, 35, 56, 7, 31, 2, 2,
	36, 40, 7, 26, 2, 2, 37, 39, 5, 14, 8, 2, 38, 37, 3, 2, 2, 2, 39, 42, 3,
	2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 46, 3, 2, 2, 2, 42,
	40, 3, 2, 2, 2, 43, 45, 5, 8, 5, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2,
	2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 52, 3, 2, 2, 2, 48, 46,
	3, 2, 2, 2, 49, 51, 5, 10, 6, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2,
	52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3,
	2, 2, 2, 55, 57, 7, 27, 2, 2, 56, 36, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	7, 3, 2, 2, 2, 58, 61, 7, 32, 2, 2, 59, 61, 7, 33, 2, 2, 60, 58, 3, 2,
	2, 2, 60, 59, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 80, 7, 37, 2, 2, 63, 67,
	7, 26, 2, 2, 64, 66, 5, 14, 8, 2, 65, 64, 3, 2, 2, 2, 66, 69, 3, 2, 2,
	2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67,
	3, 2, 2, 2, 70, 72, 7, 41, 2, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2,
	72, 76, 3, 2, 2, 2, 73, 75, 5, 12, 7, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3,
	2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3, 2, 2, 2, 78,
	76, 3, 2, 2, 2, 79, 81, 7, 27, 2, 2, 80, 63, 3, 2, 2, 2, 80, 81, 3, 2,
	2, 2, 81, 207, 3, 2, 2, 2, 82, 100, 7, 38, 2, 2, 83, 87, 7, 26, 2, 2, 84,
	86, 5, 14, 8, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2,
	2, 2, 87, 88, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 92,
	7, 41, 2, 2, 91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 96, 3, 2, 2, 2,
	93, 95, 5, 12, 7, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3,
	2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99,
	101, 7, 27, 2, 2, 100, 83, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 207,
	3, 2, 2, 2, 102, 123, 7, 40, 2, 2, 103, 107, 7, 26, 2, 2, 104, 106, 5,
	14, 8, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2,
	2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110,
	112, 7, 41, 2, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114,
	3, 2, 2, 2, 113, 115, 5, 16, 9, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2,
	2, 2, 115, 119, 3, 2, 2, 2, 116, 118, 5, 20, 11, 2, 117, 116, 3, 2, 2,
	2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120,
	122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 124, 7, 27, 2, 2, 123, 103,
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 207, 3, 2, 2, 2, 125, 146, 7, 39,
	2, 2, 126, 130, 7, 26, 2, 2, 127, 129, 5, 14, 8, 2, 128, 127, 3, 2, 2,
	2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131,
	134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 135, 7, 41, 2, 2, 134, 133,
	3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 138, 5, 16,
	9, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 142, 3, 2, 2, 2,
	139, 141, 5, 20, 11, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 145, 147, 7, 27, 2, 2, 146, 126, 3, 2, 2, 2, 146, 147, 3, 2,
	2, 2, 147, 207, 3, 2, 2, 2, 148, 149, 7, 54, 2, 2, 149, 150, 7, 26, 2,
	2, 150, 151, 5, 20, 11, 2, 151, 152, 7, 27, 2, 2, 152, 207, 3, 2, 2, 2,
	153, 154, 7, 55, 2, 2, 154, 155, 7, 26, 2, 2, 155, 156, 5, 20, 11, 2, 156,
	157, 7, 27, 2, 2, 157, 207, 3, 2, 2, 2, 158, 159, 7, 56, 2, 2, 159, 160,
	7, 26, 2, 2, 160, 161, 5, 20, 11, 2, 161, 162, 7, 27, 2, 2, 162, 207, 3,
	2, 2, 2, 163, 166, 7, 37, 2, 2, 164, 166, 7, 38, 2, 2, 165, 163, 3, 2,
	2, 2, 165, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 171, 7, 26, 2, 2,
	168, 170, 5, 14, 8, 2, 169, 168, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171,
	169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 171,
	3, 2, 2, 2, 174, 178, 7, 29, 2, 2, 175, 177, 5, 12, 7, 2, 176, 175, 3,
	2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2,
	2, 179, 181, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 207, 7, 27, 2, 2, 182,
	185, 7, 40, 2, 2, 183, 185, 7, 39, 2, 2, 184, 182, 3, 2, 2, 2, 184, 183,
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 190, 7, 26, 2, 2, 187, 189, 5, 14,
	8, 2, 188, 187, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2,
	190, 191, 3, 2, 2, 2, 191, 193, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 193,
	195, 7, 29, 2, 2, 194, 196, 5, 16, 9, 2, 195, 194, 3, 2, 2, 2, 195, 196,
	3, 2, 2, 2, 196, 200, 3, 2, 2, 2, 197, 199, 5, 20, 11, 2, 198, 197, 3,
	2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2,
	2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 205, 7, 27, 2, 2, 204,
	203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 207, 3, 2, 2, 2, 206, 62, 3,
	2, 2, 2, 206, 82, 3, 2, 2, 2, 206, 102, 3, 2, 2, 2, 206, 125, 3, 2, 2,
	2, 206, 148, 3, 2, 2, 2, 206, 153, 3, 2, 2, 2, 206, 158, 3, 2, 2, 2, 206,
	165, 3, 2, 2, 2, 206, 184, 3, 2, 2, 2, 207, 11, 3, 2, 2, 2, 208, 223, 7,
	45, 2, 2, 209, 213, 7, 26, 2, 2, 210, 212, 5, 14, 8, 2, 211, 210, 3, 2,
	2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2,
	214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 218, 5, 16, 9, 2, 217,
	216, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 220, 3, 2, 2, 2, 219, 221,
	5, 20, 11, 2, 220, 219, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3,
	2, 2, 2, 222, 224, 7, 27, 2, 2, 223, 209, 3, 2, 2, 2, 223, 224, 3, 2, 2,
	2, 224, 13, 3, 2, 2, 2, 225, 230, 7, 34, 2, 2, 226, 227, 7, 26, 2, 2, 227,
	228, 5, 20, 11, 2, 228, 229, 7, 27, 2, 2, 229, 231, 3, 2, 2, 2, 230, 226,
	3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 247, 3, 2, 2, 2, 232, 237, 7, 35,
	2, 2, 233, 234, 7, 26, 2, 2, 234, 235, 5, 20, 11, 2, 235, 236, 7, 27, 2,
	2, 236, 238, 3, 2, 2, 2, 237, 233, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	247, 3, 2, 2, 2, 239, 244, 7, 36, 2, 2, 240, 241, 7, 26, 2, 2, 241, 242,
	5, 20, 11, 2, 242, 243, 7, 27, 2, 2, 243, 245, 3, 2, 2, 2, 244, 240, 3,
	2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 225, 3, 2, 2,
	2, 246, 232, 3, 2, 2, 2, 246, 239, 3, 2, 2, 2, 247, 15, 3, 2, 2, 2, 248,
	259, 7, 42, 2, 2, 249, 250, 7, 43, 2, 2, 250, 252, 7, 26, 2, 2, 251, 253,
	5, 18, 10, 2, 252, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 252, 3,
	2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 7, 27, 2,
	2, 257, 259, 3, 2, 2, 2, 258, 248, 3, 2, 2, 2, 258, 249, 3, 2, 2, 2, 259,
	17, 3, 2, 2, 2, 260, 269, 7, 44, 2, 2, 261, 263, 7, 26, 2, 2, 262, 264,
	5, 18, 10, 2, 263, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 263, 3,
	2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268, 7, 27, 2,
	2, 268, 270, 3, 2, 2, 2, 269, 261, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270,
	19, 3, 2, 2, 2, 271, 299, 7, 47, 2, 2, 272, 299, 7, 48, 2, 2, 273, 299,
	7, 49, 2, 2, 274, 299, 7, 50, 2, 2, 275, 299, 7, 51, 2, 2, 276, 285, 7,
	52, 2, 2, 277, 279, 7, 26, 2, 2, 278, 280, 5, 20, 11, 2, 279, 278, 3, 2,
	2, 2, 280, 281, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2,
	282, 283, 3, 2, 2, 2, 283, 284, 7, 27, 2, 2, 284, 286, 3, 2, 2, 2, 285,
	277, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 299, 3, 2, 2, 2, 287, 296,
	7, 53, 2, 2, 288, 290, 7, 26, 2, 2, 289, 291, 5, 20, 11, 2, 290, 289, 3,
	2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2,
	2, 293, 294, 3, 2, 2, 2, 294, 295, 7, 27, 2, 2, 295, 297, 3, 2, 2, 2, 296,
	288, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 299, 3, 2, 2, 2, 298, 271,
	3, 2, 2, 2, 298, 272, 3, 2, 2, 2, 298, 273, 3, 2, 2, 2, 298, 274, 3, 2,
	2, 2, 298, 275, 3, 2, 2, 2, 298, 276, 3, 2, 2, 2, 298, 287, 3, 2, 2, 2,
	299, 21, 3, 2, 2, 2, 51, 40, 46, 52, 56, 60, 67, 71, 76, 80, 87, 91, 96,
	100, 107, 111, 114, 119, 123, 130, 134, 137, 142, 146, 165, 171, 178, 184,
	190, 195, 200, 204, 206, 213, 217, 220, 223, 230, 237, 244, 246, 254, 258,
	265, 269, 281, 285, 292, 296, 298,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'['", "']'", "'='", "'\"'", "'''", "';'", "'::'", "':'",
	"'.'", "','", "'<'", "'>'", "'*'", "'@'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "LSQ", "RSQ", "EQ", "DQ", "SQ", "SEMI", "DCOLON", "COLON",
	"DOT", "COMMA", "LCHEVR", "RCHEVR", "STAR", "AT", "STR", "ID", "INT", "FLT",
	"WS", "LINE_DOC", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "ADL",
	"Module", "ImportModule", "ImportScopedName", "Annotation", "AnnotationNotScoped",
	"AnnotationScoped", "Struct", "Union", "Newtype", "Type", "TypeParam",
	"TypeExprPrimOrParam", "TypeExprTypeExpr", "TypeExprElem", "Field", "Json",
	"JsonStr", "JsonBool", "JsonNull", "JsonInt", "JsonFloat", "JsonArray",
	"JsonObj", "ModuleAnno", "DeclAnno", "FieldAnno", "DNAC", "Name", "Exnotation",
}

var ruleNames = []string{
	"adl", "json", "module", "import_", "tld", "nameBody", "annotation", "typeExpr_",
	"typeExprElem_", "jsonVal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AdlWi struct {
	*antlr.BaseParser
}

func NewAdlWi(input antlr.TokenStream) *AdlWi {
	this := new(AdlWi)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AdlWi.g4"

	return this
}

// AdlWi tokens.
const (
	AdlWiEOF                 = antlr.TokenEOF
	AdlWiLCUR                = 1
	AdlWiRCUR                = 2
	AdlWiLSQ                 = 3
	AdlWiRSQ                 = 4
	AdlWiEQ                  = 5
	AdlWiDQ                  = 6
	AdlWiSQ                  = 7
	AdlWiSEMI                = 8
	AdlWiDCOLON              = 9
	AdlWiCOLON               = 10
	AdlWiDOT                 = 11
	AdlWiCOMMA               = 12
	AdlWiLCHEVR              = 13
	AdlWiRCHEVR              = 14
	AdlWiSTAR                = 15
	AdlWiAT                  = 16
	AdlWiSTR                 = 17
	AdlWiID                  = 18
	AdlWiINT                 = 19
	AdlWiFLT                 = 20
	AdlWiWS                  = 21
	AdlWiLINE_DOC            = 22
	AdlWiLINE_COMMENT        = 23
	AdlWiDOWN                = 24
	AdlWiUP                  = 25
	AdlWiROOT                = 26
	AdlWiERROR               = 27
	AdlWiADL                 = 28
	AdlWiModule              = 29
	AdlWiImportModule        = 30
	AdlWiImportScopedName    = 31
	AdlWiAnnotation          = 32
	AdlWiAnnotationNotScoped = 33
	AdlWiAnnotationScoped    = 34
	AdlWiStruct              = 35
	AdlWiUnion               = 36
	AdlWiNewtype             = 37
	AdlWiType                = 38
	AdlWiTypeParam           = 39
	AdlWiTypeExprPrimOrParam = 40
	AdlWiTypeExprTypeExpr    = 41
	AdlWiTypeExprElem        = 42
	AdlWiField               = 43
	AdlWiJson                = 44
	AdlWiJsonStr             = 45
	AdlWiJsonBool            = 46
	AdlWiJsonNull            = 47
	AdlWiJsonInt             = 48
	AdlWiJsonFloat           = 49
	AdlWiJsonArray           = 50
	AdlWiJsonObj             = 51
	AdlWiModuleAnno          = 52
	AdlWiDeclAnno            = 53
	AdlWiFieldAnno           = 54
	AdlWiDNAC                = 55
	AdlWiName                = 56
	AdlWiExnotation          = 57
)

// AdlWi rules.
const (
	AdlWiRULE_adl           = 0
	AdlWiRULE_json          = 1
	AdlWiRULE_module        = 2
	AdlWiRULE_import_       = 3
	AdlWiRULE_tld           = 4
	AdlWiRULE_nameBody      = 5
	AdlWiRULE_annotation    = 6
	AdlWiRULE_typeExpr_     = 7
	AdlWiRULE_typeExprElem_ = 8
	AdlWiRULE_jsonVal       = 9
)

type IAdlContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	Module() IModuleContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	EOF() antlr.TerminalNode
	ADL() antlr.TerminalNode
	//tokenListGetterDecl
	AllDOWN() []antlr.TerminalNode
	AllUP() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	DOWN(i int) antlr.TerminalNode
	UP(i int) antlr.TerminalNode

	// IsAdlContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *AdlContext) GetTok() antlr.Token  { return s.tok }
func (s *AdlContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AdlContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(AdlWiDOWN)
}

func (s *AdlContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, i)
}

func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(AdlWiUP)
}

func (s *AdlContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(AdlWiUP, i)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdlWiEOF, 0)
}

func (s *AdlContext) ADL() antlr.TerminalNode {
	return s.GetToken(AdlWiADL, 0)
}

//provideCopyFrom
func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlEntryListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlExitListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (s *AdlContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Adl != nil {
		h.Adl(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *AdlContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AdlContextVisitor:
		return t.VisitAdl(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWi) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdlWiRULE_adl)

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
		p.SetState(20)
		p.Match(AdlWiDOWN)
	}
	{
		p.SetState(21)
		var _m = p.Match(AdlWiADL)
		localctx.(*AdlContext).tok = _m

	}
	{
		p.SetState(22)
		p.Match(AdlWiDOWN)
	}
	{
		p.SetState(23)
		p.Module()
	}
	{
		p.SetState(24)
		p.Match(AdlWiUP)
	}
	{
		p.SetState(25)
		p.Match(AdlWiUP)
	}
	{
		p.SetState(26)
		p.Match(AdlWiEOF)
	}

	return localctx
}

type IJsonContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	EOF() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsJsonContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *JsonContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *JsonContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *JsonContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdlWiEOF, 0)
}

//provideCopyFrom
func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonEntryListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonExitListener); ok {
		listenerT.ExitJson(s)
	}
}

func (s *JsonContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Json != nil {
		h.Json(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonContextVisitor:
		return t.VisitJson(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWi) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdlWiRULE_json)

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
		p.SetState(28)
		p.Match(AdlWiDOWN)
	}
	{
		p.SetState(29)
		p.JsonVal()
	}
	{
		p.SetState(30)
		p.Match(AdlWiUP)
	}
	{
		p.SetState(31)
		p.Match(AdlWiEOF)
	}

	return localctx
}

type IModuleContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllImport_() []IImport_Context
	AllTld() []ITldContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	Import_(i int) IImport_Context
	Tld(i int) ITldContext
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Module() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsModuleContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *ModuleContext) GetTok() antlr.Token  { return s.tok }
func (s *ModuleContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *ModuleContext) Module() antlr.TerminalNode {
	return s.GetToken(AdlWiModule, 0)
}

func (s *ModuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *ModuleContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *ModuleContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ModuleContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *ModuleContext) AllImport_() []IImport_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*Import_Context)(nil)).Elem())
	var tst = make([]IImport_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_Context)
		}
	}

	return tst
}

func (s *ModuleContext) Import_(i int) IImport_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Import_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *ModuleContext) AllTld() []ITldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *ModuleContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
}

//provideCopyFrom
func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleEntryListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleExitListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Module != nil {
		h.Module(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleContextVisitor:
		return t.VisitModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWi) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AdlWiRULE_module)
	var //TokenTypeDecl
	_la int

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
		p.SetState(33)
		var _m = p.Match(AdlWiModule)
		localctx.(*ModuleContext).tok = _m

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWiDOWN {
		{
			p.SetState(34)
			p.Match(AdlWiDOWN)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(35)
				p.Annotation()
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlWiImportModule || _la == AdlWiImportScopedName {
			{
				p.SetState(41)
				p.Import_()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(AdlWiStruct-35))|(1<<(AdlWiUnion-35))|(1<<(AdlWiNewtype-35))|(1<<(AdlWiType-35))|(1<<(AdlWiModuleAnno-35))|(1<<(AdlWiDeclAnno-35))|(1<<(AdlWiFieldAnno-35)))) != 0 {
			{
				p.SetState(47)
				p.Tld()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(53)
			p.Match(AdlWiUP)
		}

	}

	return localctx
}

type IImport_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsImport_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Import_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_import_
	return p
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *Import_Context) CopyFrom(ctx *Import_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ImportScopedModuleContext struct {
	*Import_Context
}

func NewImportScopedModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportScopedModuleContext {
	var p = new(ImportScopedModuleContext)

	p.Import_Context = NewEmptyImport_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Import_Context))

	return p
}

type IImportScopedModuleContext interface {
	//Current rule
	IImport_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ImportScopedName() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ImportScopedModuleContext) IsImportScopedModuleContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportScopedModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportScopedModuleContext) ImportScopedName() antlr.TerminalNode {
	return s.GetToken(AdlWiImportScopedName, 0)
}

func (s *ImportScopedModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedModuleEntryListener); ok {
		listenerT.EnterImportScopedModule(s)
	}
}

func (s *ImportScopedModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedModuleExitListener); ok {
		listenerT.ExitImportScopedModule(s)
	}
}

func (s *ImportScopedModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportScopedModule != nil {
		h.ImportScopedModule(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportScopedModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportScopedModuleContextVisitor:
		return t.VisitImportScopedModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ImportModuleContext struct {
	*Import_Context
}

func NewImportModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportModuleContext {
	var p = new(ImportModuleContext)

	p.Import_Context = NewEmptyImport_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Import_Context))

	return p
}

type IImportModuleContext interface {
	//Current rule
	IImport_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ImportModule() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ImportModuleContext) IsImportModuleContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportModuleContext) ImportModule() antlr.TerminalNode {
	return s.GetToken(AdlWiImportModule, 0)
}

func (s *ImportModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleEntryListener); ok {
		listenerT.EnterImportModule(s)
	}
}

func (s *ImportModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleExitListener); ok {
		listenerT.ExitImportModule(s)
	}
}

func (s *ImportModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportModule != nil {
		h.ImportModule(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportModuleContextVisitor:
		return t.VisitImportModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWi) Import_() (localctx IImport_Context) {
	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AdlWiRULE_import_)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiImportModule:
		localctx = NewImportModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(AdlWiImportModule)
		}

	case AdlWiImportScopedName:
		localctx = NewImportScopedModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(AdlWiImportScopedName)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITldContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTldContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTldContext() *TldContext {
	var p = new(TldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_tld

	return p
}

func (s *TldContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TldContext) CopyFrom(ctx *TldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type ITypeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	Type() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeContext) IsTypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeContext) GetTok() antlr.Token  { return s.tok }
func (s *TypeContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeContext) Type() antlr.TerminalNode {
	return s.GetToken(AdlWiType, 0)
}

func (s *TypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *TypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *TypeContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeParam, 0)
}

func (s *TypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeEntryListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExitListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Type != nil {
		h.Type(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeContextVisitor:
		return t.VisitType(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type TypeParamErrorContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewTypeParamErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type ITypeParamErrorContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllNameBody() []INameBodyContext
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	NameBody(i int) INameBodyContext
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	ERROR() antlr.TerminalNode
	UP() antlr.TerminalNode
	Struct() antlr.TerminalNode
	Union() antlr.TerminalNode
	Type() antlr.TerminalNode
	Newtype() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeParamErrorContext) IsTypeParamErrorContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeParamErrorContext) GetTok() antlr.Token  { return s.tok }
func (s *TypeParamErrorContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeParamErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeParamErrorContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *TypeParamErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlWiERROR, 0)
}

func (s *TypeParamErrorContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *TypeParamErrorContext) Struct() antlr.TerminalNode {
	return s.GetToken(AdlWiStruct, 0)
}

func (s *TypeParamErrorContext) Union() antlr.TerminalNode {
	return s.GetToken(AdlWiUnion, 0)
}

func (s *TypeParamErrorContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TypeParamErrorContext) AllNameBody() []INameBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*NameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *TypeParamErrorContext) Type() antlr.TerminalNode {
	return s.GetToken(AdlWiType, 0)
}

func (s *TypeParamErrorContext) Newtype() antlr.TerminalNode {
	return s.GetToken(AdlWiNewtype, 0)
}

func (s *TypeParamErrorContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeParamErrorContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeParamErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorEntryListener); ok {
		listenerT.EnterTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorExitListener); ok {
		listenerT.ExitTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParamError != nil {
		h.TypeParamError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParamErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParamErrorContextVisitor:
		return t.VisitTypeParamError(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type NewtypeContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewtypeContext {
	var p = new(NewtypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type INewtypeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	Newtype() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NewtypeContext) IsNewtypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *NewtypeContext) GetTok() antlr.Token  { return s.tok }
func (s *NewtypeContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NewtypeContext) Newtype() antlr.TerminalNode {
	return s.GetToken(AdlWiNewtype, 0)
}

func (s *NewtypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *NewtypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *NewtypeContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NewtypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NewtypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeParam, 0)
}

func (s *NewtypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *NewtypeContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *NewtypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewtypeEntryListener); ok {
		listenerT.EnterNewtype(s)
	}
}

func (s *NewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewtypeExitListener); ok {
		listenerT.ExitNewtype(s)
	}
}

func (s *NewtypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Newtype != nil {
		h.Newtype(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NewtypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NewtypeContextVisitor:
		return t.VisitNewtype(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ModAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewModAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAnnoContext {
	var p = new(ModAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IModAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	ModuleAnno() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ModAnnoContext) IsModAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *ModAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *ModAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *ModAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *ModAnnoContext) ModuleAnno() antlr.TerminalNode {
	return s.GetToken(AdlWiModuleAnno, 0)
}

func (s *ModAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModAnnoEntryListener); ok {
		listenerT.EnterModAnno(s)
	}
}

func (s *ModAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModAnnoExitListener); ok {
		listenerT.ExitModAnno(s)
	}
}

func (s *ModAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModAnno != nil {
		h.ModAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModAnnoContextVisitor:
		return t.VisitModAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FieldAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewFieldAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnoContext {
	var p = new(FieldAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IFieldAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	FieldAnno() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldAnnoContext) IsFieldAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *FieldAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *FieldAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *FieldAnnoContext) FieldAnno() antlr.TerminalNode {
	return s.GetToken(AdlWiFieldAnno, 0)
}

func (s *FieldAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnoEntryListener); ok {
		listenerT.EnterFieldAnno(s)
	}
}

func (s *FieldAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnoExitListener); ok {
		listenerT.ExitFieldAnno(s)
	}
}

func (s *FieldAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldAnno != nil {
		h.FieldAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldAnnoContextVisitor:
		return t.VisitFieldAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type UnionContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IUnionContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllNameBody() []INameBodyContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	NameBody(i int) INameBodyContext

	//  tokenGetterDecl
	Union() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*UnionContext) IsUnionContext() {}

//AltLabelStructDecl tokenDecls
func (s *UnionContext) GetTok() antlr.Token  { return s.tok }
func (s *UnionContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *UnionContext) Union() antlr.TerminalNode {
	return s.GetToken(AdlWiUnion, 0)
}

func (s *UnionContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *UnionContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *UnionContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *UnionContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *UnionContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeParam, 0)
}

func (s *UnionContext) AllNameBody() []INameBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*NameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *UnionContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnionEntryListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnionExitListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Union != nil {
		h.Union(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *UnionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case UnionContextVisitor:
		return t.VisitUnion(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type DeclAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewDeclAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnoContext {
	var p = new(DeclAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IDeclAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	DeclAnno() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*DeclAnnoContext) IsDeclAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *DeclAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *DeclAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DeclAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DeclAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *DeclAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *DeclAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *DeclAnnoContext) DeclAnno() antlr.TerminalNode {
	return s.GetToken(AdlWiDeclAnno, 0)
}

func (s *DeclAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnoEntryListener); ok {
		listenerT.EnterDeclAnno(s)
	}
}

func (s *DeclAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnoExitListener); ok {
		listenerT.ExitDeclAnno(s)
	}
}

func (s *DeclAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DeclAnno != nil {
		h.DeclAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclAnnoContextVisitor:
		return t.VisitDeclAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type StructContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IStructContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllNameBody() []INameBodyContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	NameBody(i int) INameBodyContext

	//  tokenGetterDecl
	Struct() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*StructContext) IsStructContext() {}

//AltLabelStructDecl tokenDecls
func (s *StructContext) GetTok() antlr.Token  { return s.tok }
func (s *StructContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StructContext) Struct() antlr.TerminalNode {
	return s.GetToken(AdlWiStruct, 0)
}

func (s *StructContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *StructContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *StructContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StructContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeParam, 0)
}

func (s *StructContext) AllNameBody() []INameBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*NameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *StructContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructEntryListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructExitListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Struct != nil {
		h.Struct(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StructContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StructContextVisitor:
		return t.VisitStruct(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWi) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AdlWiRULE_tld)
	var //TokenTypeDecl
	_la int

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			var _m = p.Match(AdlWiStruct)
			localctx.(*StructContext).tok = _m

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(61)
				p.Match(AdlWiDOWN)
			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(62)
					p.Annotation()
				}

				p.SetState(67)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(68)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWiField {
				{
					p.SetState(71)
					p.NameBody()
				}

				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(77)
				p.Match(AdlWiUP)
			}

		}

	case 2:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			var _m = p.Match(AdlWiUnion)
			localctx.(*UnionContext).tok = _m

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(81)
				p.Match(AdlWiDOWN)
			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(82)
					p.Annotation()
				}

				p.SetState(87)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(88)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWiField {
				{
					p.SetState(91)
					p.NameBody()
				}

				p.SetState(96)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(97)
				p.Match(AdlWiUP)
			}

		}

	case 3:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			var _m = p.Match(AdlWiType)
			localctx.(*TypeContext).tok = _m

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(101)
				p.Match(AdlWiDOWN)
			}
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(102)
					p.Annotation()
				}

				p.SetState(107)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(108)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
				{
					p.SetState(111)
					p.TypeExpr_()
				}

			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0 {
				{
					p.SetState(114)
					p.JsonVal()
				}

				p.SetState(119)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(120)
				p.Match(AdlWiUP)
			}

		}

	case 4:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			var _m = p.Match(AdlWiNewtype)
			localctx.(*NewtypeContext).tok = _m

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(124)
				p.Match(AdlWiDOWN)
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(125)
					p.Annotation()
				}

				p.SetState(130)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(131)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
				{
					p.SetState(134)
					p.TypeExpr_()
				}

			}
			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0 {
				{
					p.SetState(137)
					p.JsonVal()
				}

				p.SetState(142)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(143)
				p.Match(AdlWiUP)
			}

		}

	case 5:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			var _m = p.Match(AdlWiModuleAnno)
			localctx.(*ModAnnoContext).tok = _m

		}
		{
			p.SetState(147)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(148)
			p.JsonVal()
		}
		{
			p.SetState(149)
			p.Match(AdlWiUP)
		}

	case 6:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
			var _m = p.Match(AdlWiDeclAnno)
			localctx.(*DeclAnnoContext).tok = _m

		}
		{
			p.SetState(152)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(153)
			p.JsonVal()
		}
		{
			p.SetState(154)
			p.Match(AdlWiUP)
		}

	case 7:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(156)
			var _m = p.Match(AdlWiFieldAnno)
			localctx.(*FieldAnnoContext).tok = _m

		}
		{
			p.SetState(157)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(158)
			p.JsonVal()
		}
		{
			p.SetState(159)
			p.Match(AdlWiUP)
		}

	case 8:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(163)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AdlWiStruct:
			{
				p.SetState(161)
				var _m = p.Match(AdlWiStruct)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		case AdlWiUnion:
			{
				p.SetState(162)
				var _m = p.Match(AdlWiUnion)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(165)
			p.Match(AdlWiDOWN)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(166)
				p.Annotation()
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(172)
			p.Match(AdlWiERROR)
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlWiField {
			{
				p.SetState(173)
				p.NameBody()
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(179)
			p.Match(AdlWiUP)
		}

	case 9:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(182)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AdlWiType:
			{
				p.SetState(180)
				var _m = p.Match(AdlWiType)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		case AdlWiNewtype:
			{
				p.SetState(181)
				var _m = p.Match(AdlWiNewtype)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(184)
			p.Match(AdlWiDOWN)
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(185)
				p.Annotation()
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(191)
			p.Match(AdlWiERROR)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
			{
				p.SetState(192)
				p.TypeExpr_()
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0 {
			{
				p.SetState(195)
				p.JsonVal()
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(201)
				p.Match(AdlWiUP)
			}

		}

	}

	return localctx
}

type INameBodyContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsNameBodyContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type NameBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameBodyContext() *NameBodyContext {
	var p = new(NameBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_nameBody
	return p
}

func (*NameBodyContext) IsNameBodyContext() {}

func NewNameBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameBodyContext {
	var p = new(NameBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_nameBody

	return p
}

func (s *NameBodyContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *NameBodyContext) CopyFrom(ctx *NameBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type FieldContext struct {
	*NameBodyContext
	//TokenDecl
	tok antlr.Token
}

func NewFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldContext {
	var p = new(FieldContext)

	p.NameBodyContext = NewEmptyNameBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameBodyContext))

	return p
}

type IFieldContext interface {
	//Current rule
	INameBodyContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext

	//  tokenGetterDecl
	Field() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldContext) IsFieldContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldContext) GetTok() antlr.Token  { return s.tok }
func (s *FieldContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AdlWiField, 0)
}

func (s *FieldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *FieldContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *FieldContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FieldContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FieldContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *FieldContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldEntryListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldExitListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Field != nil {
		h.Field(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldContextVisitor:
		return t.VisitField(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWi) NameBody() (localctx INameBodyContext) {
	localctx = NewNameBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AdlWiRULE_nameBody)
	var //TokenTypeDecl
	_la int

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

	localctx = NewFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		var _m = p.Match(AdlWiField)
		localctx.(*FieldContext).tok = _m

	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWiDOWN {
		{
			p.SetState(207)
			p.Match(AdlWiDOWN)
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(208)
				p.Annotation()
			}

			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
			{
				p.SetState(214)
				p.TypeExpr_()
			}

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0 {
			{
				p.SetState(217)
				p.JsonVal()
			}

		}
		{
			p.SetState(220)
			p.Match(AdlWiUP)
		}

	}

	return localctx
}

type IAnnotationContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Annotation() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	AnnotationNotScoped() antlr.TerminalNode
	AnnotationScoped() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsAnnotationContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *AnnotationContext) GetTok() antlr.Token  { return s.tok }
func (s *AnnotationContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AnnotationContext) Annotation() antlr.TerminalNode {
	return s.GetToken(AdlWiAnnotation, 0)
}

func (s *AnnotationContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *AnnotationContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *AnnotationContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *AnnotationContext) AnnotationNotScoped() antlr.TerminalNode {
	return s.GetToken(AdlWiAnnotationNotScoped, 0)
}

func (s *AnnotationContext) AnnotationScoped() antlr.TerminalNode {
	return s.GetToken(AdlWiAnnotationScoped, 0)
}

//provideCopyFrom
func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationEntryListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationExitListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Annotation != nil {
		h.Annotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *AnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AnnotationContextVisitor:
		return t.VisitAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWi) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AdlWiRULE_annotation)
	var //TokenTypeDecl
	_la int

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

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiAnnotation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			var _m = p.Match(AdlWiAnnotation)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(224)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(225)
				p.JsonVal()
			}
			{
				p.SetState(226)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiAnnotationNotScoped:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			var _m = p.Match(AdlWiAnnotationNotScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(231)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(232)
				p.JsonVal()
			}
			{
				p.SetState(233)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiAnnotationScoped:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			var _m = p.Match(AdlWiAnnotationScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(238)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(239)
				p.JsonVal()
			}
			{
				p.SetState(240)
				p.Match(AdlWiUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITypeExpr_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem_() []ITypeExprElem_Context
	//  ruleListIndexedGetterDecl
	TypeExprElem_(i int) ITypeExprElem_Context
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	TypeExprPrimOrParam() antlr.TerminalNode
	TypeExprTypeExpr() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExpr_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExpr_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyTypeExpr_Context() *TypeExpr_Context {
	var p = new(TypeExpr_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_typeExpr_
	return p
}

func (*TypeExpr_Context) IsTypeExpr_Context() {}

func NewTypeExpr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExpr_Context {
	var p = new(TypeExpr_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_typeExpr_

	return p
}

func (s *TypeExpr_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *TypeExpr_Context) GetTok() antlr.Token  { return s.tok }
func (s *TypeExpr_Context) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeExpr_Context) TypeExprPrimOrParam() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeExprPrimOrParam, 0)
}

func (s *TypeExpr_Context) TypeExprTypeExpr() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeExprTypeExpr, 0)
}

func (s *TypeExpr_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *TypeExpr_Context) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *TypeExpr_Context) AllTypeExprElem_() []ITypeExprElem_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeExpr_Context) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

//provideCopyFrom
func (s *TypeExpr_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpr_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeExpr_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpr_EntryListener); ok {
		listenerT.EnterTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpr_ExitListener); ok {
		listenerT.ExitTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpr_ != nil {
		h.TypeExpr_(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpr_Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpr_ContextVisitor:
		return t.VisitTypeExpr_(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWi) TypeExpr_() (localctx ITypeExpr_Context) {
	localctx = NewTypeExpr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AdlWiRULE_typeExpr_)
	var //TokenTypeDecl
	_la int

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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiTypeExprPrimOrParam:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			var _m = p.Match(AdlWiTypeExprPrimOrParam)
			localctx.(*TypeExpr_Context).tok = _m

		}

	case AdlWiTypeExprTypeExpr:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			var _m = p.Match(AdlWiTypeExprTypeExpr)
			localctx.(*TypeExpr_Context).tok = _m

		}

		{
			p.SetState(248)
			p.Match(AdlWiDOWN)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWiTypeExprElem {
			{
				p.SetState(249)
				p.TypeExprElem_()
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(254)
			p.Match(AdlWiUP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITypeExprElem_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExprElem_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprElem_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElem_Context() *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_typeExprElem_
	return p
}

func (*TypeExprElem_Context) IsTypeExprElem_Context() {}

func NewTypeExprElem_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_typeExprElem_

	return p
}

func (s *TypeExprElem_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprElem_Context) CopyFrom(ctx *TypeExprElem_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElem_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElem_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeParamsContext struct {
	*TypeExprElem_Context
	//TokenDecl
	tok antlr.Token
}

func NewTypeParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamsContext {
	var p = new(TypeParamsContext)

	p.TypeExprElem_Context = NewEmptyTypeExprElem_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElem_Context))

	return p
}

type ITypeParamsContext interface {
	//Current rule
	ITypeExprElem_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem_() []ITypeExprElem_Context
	//  ruleListIndexedGetterDecl
	TypeExprElem_(i int) ITypeExprElem_Context

	//  tokenGetterDecl
	TypeExprElem() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeParamsContext) IsTypeParamsContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeParamsContext) GetTok() antlr.Token  { return s.tok }
func (s *TypeParamsContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeParamsContext) TypeExprElem() antlr.TerminalNode {
	return s.GetToken(AdlWiTypeExprElem, 0)
}

func (s *TypeParamsContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *TypeParamsContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *TypeParamsContext) AllTypeExprElem_() []ITypeExprElem_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeParamsContext) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamsEntryListener); ok {
		listenerT.EnterTypeParams(s)
	}
}

func (s *TypeParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamsExitListener); ok {
		listenerT.ExitTypeParams(s)
	}
}

func (s *TypeParamsContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParams != nil {
		h.TypeParams(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParamsContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParamsContextVisitor:
		return t.VisitTypeParams(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWi) TypeExprElem_() (localctx ITypeExprElem_Context) {
	localctx = NewTypeExprElem_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AdlWiRULE_typeExprElem_)
	var //TokenTypeDecl
	_la int

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

	localctx = NewTypeParamsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		var _m = p.Match(AdlWiTypeExprElem)
		localctx.(*TypeParamsContext).tok = _m

	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWiDOWN {
		{
			p.SetState(259)
			p.Match(AdlWiDOWN)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWiTypeExprElem {
			{
				p.SetState(260)
				p.TypeExprElem_()
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(265)
			p.Match(AdlWiUP)
		}

	}

	return localctx
}

type IJsonValContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsJsonValContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValContext() *JsonValContext {
	var p = new(JsonValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWiRULE_jsonVal
	return p
}

func (*JsonValContext) IsJsonValContext() {}

func NewJsonValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValContext {
	var p = new(JsonValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWiRULE_jsonVal

	return p
}

func (s *JsonValContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *JsonValContext) CopyFrom(ctx *JsonValContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type JsonStrContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStrContext {
	var p = new(JsonStrContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonStrContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonStr() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonStrContext) IsJsonStrContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonStrContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonStrContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonStr, 0)
}

func (s *JsonStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonStrEntryListener); ok {
		listenerT.EnterJsonStr(s)
	}
}

func (s *JsonStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonStrExitListener); ok {
		listenerT.ExitJsonStr(s)
	}
}

func (s *JsonStrContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonStr != nil {
		h.JsonStr(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonStrContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonStrContextVisitor:
		return t.VisitJsonStr(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonArrayContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonArrayContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	JsonArray() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonArrayContext) IsJsonArrayContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonArrayContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonArrayContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonArray, 0)
}

func (s *JsonArrayContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *JsonArrayContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *JsonArrayContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonArrayContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonArrayEntryListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonArrayExitListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

func (s *JsonArrayContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonArray != nil {
		h.JsonArray(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonArrayContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonArrayContextVisitor:
		return t.VisitJsonArray(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonFloatContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonFloatContext {
	var p = new(JsonFloatContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonFloatContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonFloat() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonFloatContext) IsJsonFloatContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonFloatContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonFloatContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonFloat, 0)
}

func (s *JsonFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonFloatEntryListener); ok {
		listenerT.EnterJsonFloat(s)
	}
}

func (s *JsonFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonFloatExitListener); ok {
		listenerT.ExitJsonFloat(s)
	}
}

func (s *JsonFloatContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonFloat != nil {
		h.JsonFloat(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonFloatContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonFloatContextVisitor:
		return t.VisitJsonFloat(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonObjContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjContext {
	var p = new(JsonObjContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonObjContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	JsonObj() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonObjContext) IsJsonObjContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonObjContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonObjContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonObj, 0)
}

func (s *JsonObjContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *JsonObjContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *JsonObjContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonObjContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonObjEntryListener); ok {
		listenerT.EnterJsonObj(s)
	}
}

func (s *JsonObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonObjExitListener); ok {
		listenerT.ExitJsonObj(s)
	}
}

func (s *JsonObjContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonObj != nil {
		h.JsonObj(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonObjContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonObjContextVisitor:
		return t.VisitJsonObj(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonBoolContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonBoolContext {
	var p = new(JsonBoolContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonBoolContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonBool() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonBoolContext) IsJsonBoolContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonBoolContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonBoolContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonBool, 0)
}

func (s *JsonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonBoolEntryListener); ok {
		listenerT.EnterJsonBool(s)
	}
}

func (s *JsonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonBoolExitListener); ok {
		listenerT.ExitJsonBool(s)
	}
}

func (s *JsonBoolContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonBool != nil {
		h.JsonBool(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonBoolContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonBoolContextVisitor:
		return t.VisitJsonBool(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonIntContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonIntContext {
	var p = new(JsonIntContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonIntContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonInt() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonIntContext) IsJsonIntContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonIntContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonIntContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonInt, 0)
}

func (s *JsonIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonIntEntryListener); ok {
		listenerT.EnterJsonInt(s)
	}
}

func (s *JsonIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonIntExitListener); ok {
		listenerT.ExitJsonInt(s)
	}
}

func (s *JsonIntContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonInt != nil {
		h.JsonInt(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonIntContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonIntContextVisitor:
		return t.VisitJsonInt(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonNullContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNullContext {
	var p = new(JsonNullContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonNullContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonNull() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonNullContext) IsJsonNullContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonNullContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonNullContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(AdlWiJsonNull, 0)
}

func (s *JsonNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonNullEntryListener); ok {
		listenerT.EnterJsonNull(s)
	}
}

func (s *JsonNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonNullExitListener); ok {
		listenerT.ExitJsonNull(s)
	}
}

func (s *JsonNullContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonNull != nil {
		h.JsonNull(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonNullContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonNullContextVisitor:
		return t.VisitJsonNull(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWi) JsonVal() (localctx IJsonValContext) {
	localctx = NewJsonValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AdlWiRULE_jsonVal)
	var //TokenTypeDecl
	_la int

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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			var _m = p.Match(AdlWiJsonStr)
			localctx.(*JsonStrContext).tok = _m

		}

	case AdlWiJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			var _m = p.Match(AdlWiJsonBool)
			localctx.(*JsonBoolContext).tok = _m

		}

	case AdlWiJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(271)
			var _m = p.Match(AdlWiJsonNull)
			localctx.(*JsonNullContext).tok = _m

		}

	case AdlWiJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(272)
			var _m = p.Match(AdlWiJsonInt)
			localctx.(*JsonIntContext).tok = _m

		}

	case AdlWiJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(273)
			var _m = p.Match(AdlWiJsonFloat)
			localctx.(*JsonFloatContext).tok = _m

		}

	case AdlWiJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(274)
			var _m = p.Match(AdlWiJsonArray)
			localctx.(*JsonArrayContext).tok = _m

		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(275)
				p.Match(AdlWiDOWN)
			}
			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0) {
				{
					p.SetState(276)
					p.JsonVal()
				}

				p.SetState(279)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(281)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(285)
			var _m = p.Match(AdlWiJsonObj)
			localctx.(*JsonObjContext).tok = _m

		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(286)
				p.Match(AdlWiDOWN)
			}
			p.SetState(288)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWiJsonStr-45))|(1<<(AdlWiJsonBool-45))|(1<<(AdlWiJsonNull-45))|(1<<(AdlWiJsonInt-45))|(1<<(AdlWiJsonFloat-45))|(1<<(AdlWiJsonArray-45))|(1<<(AdlWiJsonObj-45)))) != 0) {
				{
					p.SetState(287)
					p.JsonVal()
				}

				p.SetState(290)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(292)
				p.Match(AdlWiUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
