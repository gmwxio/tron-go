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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 325,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 4, 7, 4, 45, 10, 4, 12,
	4, 14, 4, 48, 11, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 3,
	4, 5, 4, 57, 10, 4, 3, 4, 3, 4, 3, 4, 6, 4, 62, 10, 4, 13, 4, 14, 4, 63,
	3, 4, 5, 4, 67, 10, 4, 5, 4, 69, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 74, 10,
	5, 3, 6, 3, 6, 3, 6, 7, 6, 79, 10, 6, 12, 6, 14, 6, 82, 11, 6, 3, 6, 5,
	6, 85, 10, 6, 3, 6, 7, 6, 88, 10, 6, 12, 6, 14, 6, 91, 11, 6, 3, 6, 5,
	6, 94, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 99, 10, 6, 12, 6, 14, 6, 102, 11,
	6, 3, 6, 5, 6, 105, 10, 6, 3, 6, 7, 6, 108, 10, 6, 12, 6, 14, 6, 111, 11,
	6, 3, 6, 5, 6, 114, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 119, 10, 6, 12, 6, 14,
	6, 122, 11, 6, 3, 6, 5, 6, 125, 10, 6, 3, 6, 5, 6, 128, 10, 6, 3, 6, 7,
	6, 131, 10, 6, 12, 6, 14, 6, 134, 11, 6, 3, 6, 5, 6, 137, 10, 6, 3, 6,
	3, 6, 3, 6, 7, 6, 142, 10, 6, 12, 6, 14, 6, 145, 11, 6, 3, 6, 5, 6, 148,
	10, 6, 3, 6, 5, 6, 151, 10, 6, 3, 6, 7, 6, 154, 10, 6, 12, 6, 14, 6, 157,
	11, 6, 3, 6, 5, 6, 160, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 179,
	10, 6, 3, 6, 3, 6, 7, 6, 183, 10, 6, 12, 6, 14, 6, 186, 11, 6, 3, 6, 3,
	6, 7, 6, 190, 10, 6, 12, 6, 14, 6, 193, 11, 6, 3, 6, 3, 6, 3, 6, 5, 6,
	198, 10, 6, 3, 6, 3, 6, 7, 6, 202, 10, 6, 12, 6, 14, 6, 205, 11, 6, 3,
	6, 3, 6, 5, 6, 209, 10, 6, 3, 6, 7, 6, 212, 10, 6, 12, 6, 14, 6, 215, 11,
	6, 3, 6, 5, 6, 218, 10, 6, 3, 6, 3, 6, 3, 6, 6, 6, 223, 10, 6, 13, 6, 14,
	6, 224, 3, 6, 5, 6, 228, 10, 6, 5, 6, 230, 10, 6, 3, 7, 3, 7, 3, 7, 7,
	7, 235, 10, 7, 12, 7, 14, 7, 238, 11, 7, 3, 7, 5, 7, 241, 10, 7, 3, 7,
	5, 7, 244, 10, 7, 3, 7, 5, 7, 247, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	5, 8, 254, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 261, 10, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 268, 10, 8, 5, 8, 270, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 6, 9, 276, 10, 9, 13, 9, 14, 9, 277, 3, 9, 3, 9, 5, 9, 282,
	10, 9, 3, 10, 3, 10, 3, 10, 6, 10, 287, 10, 10, 13, 10, 14, 10, 288, 3,
	10, 3, 10, 5, 10, 293, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 6, 11, 303, 10, 11, 13, 11, 14, 11, 304, 3, 11, 3, 11, 5,
	11, 309, 10, 11, 3, 11, 3, 11, 3, 11, 6, 11, 314, 10, 11, 13, 11, 14, 11,
	315, 3, 11, 3, 11, 5, 11, 320, 10, 11, 3, 11, 5, 11, 323, 10, 11, 3, 11,
	4, 63, 224, 2, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 384, 2,
	22, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 73, 3, 2, 2, 2,
	10, 229, 3, 2, 2, 2, 12, 231, 3, 2, 2, 2, 14, 269, 3, 2, 2, 2, 16, 281,
	3, 2, 2, 2, 18, 283, 3, 2, 2, 2, 20, 322, 3, 2, 2, 2, 22, 23, 7, 26, 2,
	2, 23, 24, 7, 30, 2, 2, 24, 25, 7, 26, 2, 2, 25, 26, 5, 6, 4, 2, 26, 27,
	7, 27, 2, 2, 27, 28, 7, 27, 2, 2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2,
	30, 31, 7, 26, 2, 2, 31, 32, 5, 20, 11, 2, 32, 33, 7, 27, 2, 2, 33, 34,
	7, 2, 2, 3, 34, 5, 3, 2, 2, 2, 35, 56, 7, 31, 2, 2, 36, 40, 7, 26, 2, 2,
	37, 39, 5, 14, 8, 2, 38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3,
	2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 46, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43,
	45, 5, 8, 5, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2,
	2, 46, 47, 3, 2, 2, 2, 47, 52, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51,
	5, 10, 6, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2,
	52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 7,
	27, 2, 2, 56, 36, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 69, 3, 2, 2, 2, 58,
	66, 7, 29, 2, 2, 59, 61, 7, 26, 2, 2, 60, 62, 11, 2, 2, 2, 61, 60, 3, 2,
	2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65,
	3, 2, 2, 2, 65, 67, 7, 27, 2, 2, 66, 59, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 69, 3, 2, 2, 2, 68, 35, 3, 2, 2, 2, 68, 58, 3, 2, 2, 2, 69, 7, 3, 2,
	2, 2, 70, 74, 7, 32, 2, 2, 71, 74, 7, 33, 2, 2, 72, 74, 7, 29, 2, 2, 73,
	70, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 9, 3, 2, 2,
	2, 75, 93, 7, 37, 2, 2, 76, 80, 7, 26, 2, 2, 77, 79, 5, 14, 8, 2, 78, 77,
	3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2,
	81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 83, 85, 7, 41, 2, 2, 84, 83, 3,
	2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 89, 3, 2, 2, 2, 86, 88, 5, 12, 7, 2, 87,
	86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 94, 7, 27, 2, 2, 93, 76,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 230, 3, 2, 2, 2, 95, 113, 7, 38, 2,
	2, 96, 100, 7, 26, 2, 2, 97, 99, 5, 14, 8, 2, 98, 97, 3, 2, 2, 2, 99, 102,
	3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 104, 3, 2,
	2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 7, 41, 2, 2, 104, 103, 3, 2, 2, 2,
	104, 105, 3, 2, 2, 2, 105, 109, 3, 2, 2, 2, 106, 108, 5, 12, 7, 2, 107,
	106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110,
	3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 112, 114, 7, 27,
	2, 2, 113, 96, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 230, 3, 2, 2, 2,
	115, 136, 7, 40, 2, 2, 116, 120, 7, 26, 2, 2, 117, 119, 5, 14, 8, 2, 118,
	117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121,
	3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 125, 7, 41,
	2, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2,
	126, 128, 5, 16, 9, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	132, 3, 2, 2, 2, 129, 131, 5, 20, 11, 2, 130, 129, 3, 2, 2, 2, 131, 134,
	3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2,
	2, 2, 134, 132, 3, 2, 2, 2, 135, 137, 7, 27, 2, 2, 136, 116, 3, 2, 2, 2,
	136, 137, 3, 2, 2, 2, 137, 230, 3, 2, 2, 2, 138, 159, 7, 39, 2, 2, 139,
	143, 7, 26, 2, 2, 140, 142, 5, 14, 8, 2, 141, 140, 3, 2, 2, 2, 142, 145,
	3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 147, 3, 2,
	2, 2, 145, 143, 3, 2, 2, 2, 146, 148, 7, 41, 2, 2, 147, 146, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 150, 3, 2, 2, 2, 149, 151, 5, 16, 9, 2, 150,
	149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 155, 3, 2, 2, 2, 152, 154,
	5, 20, 11, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3,
	2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157, 155, 3, 2, 2,
	2, 158, 160, 7, 27, 2, 2, 159, 139, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	230, 3, 2, 2, 2, 161, 162, 7, 54, 2, 2, 162, 163, 7, 26, 2, 2, 163, 164,
	5, 20, 11, 2, 164, 165, 7, 27, 2, 2, 165, 230, 3, 2, 2, 2, 166, 167, 7,
	55, 2, 2, 167, 168, 7, 26, 2, 2, 168, 169, 5, 20, 11, 2, 169, 170, 7, 27,
	2, 2, 170, 230, 3, 2, 2, 2, 171, 172, 7, 56, 2, 2, 172, 173, 7, 26, 2,
	2, 173, 174, 5, 20, 11, 2, 174, 175, 7, 27, 2, 2, 175, 230, 3, 2, 2, 2,
	176, 179, 7, 37, 2, 2, 177, 179, 7, 38, 2, 2, 178, 176, 3, 2, 2, 2, 178,
	177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 184, 7, 26, 2, 2, 181, 183,
	5, 14, 8, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2,
	2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2,
	187, 191, 7, 29, 2, 2, 188, 190, 5, 12, 7, 2, 189, 188, 3, 2, 2, 2, 190,
	193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194,
	3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 230, 7, 27, 2, 2, 195, 198, 7, 40,
	2, 2, 196, 198, 7, 39, 2, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2, 2, 2,
	198, 199, 3, 2, 2, 2, 199, 203, 7, 26, 2, 2, 200, 202, 5, 14, 8, 2, 201,
	200, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204,
	3, 2, 2, 2, 204, 206, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 208, 7, 29,
	2, 2, 207, 209, 5, 16, 9, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2,
	209, 213, 3, 2, 2, 2, 210, 212, 5, 20, 11, 2, 211, 210, 3, 2, 2, 2, 212,
	215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 217,
	3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 218, 7, 27, 2, 2, 217, 216, 3, 2,
	2, 2, 217, 218, 3, 2, 2, 2, 218, 230, 3, 2, 2, 2, 219, 227, 7, 29, 2, 2,
	220, 222, 7, 26, 2, 2, 221, 223, 11, 2, 2, 2, 222, 221, 3, 2, 2, 2, 223,
	224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 226,
	3, 2, 2, 2, 226, 228, 7, 27, 2, 2, 227, 220, 3, 2, 2, 2, 227, 228, 3, 2,
	2, 2, 228, 230, 3, 2, 2, 2, 229, 75, 3, 2, 2, 2, 229, 95, 3, 2, 2, 2, 229,
	115, 3, 2, 2, 2, 229, 138, 3, 2, 2, 2, 229, 161, 3, 2, 2, 2, 229, 166,
	3, 2, 2, 2, 229, 171, 3, 2, 2, 2, 229, 178, 3, 2, 2, 2, 229, 197, 3, 2,
	2, 2, 229, 219, 3, 2, 2, 2, 230, 11, 3, 2, 2, 2, 231, 246, 7, 45, 2, 2,
	232, 236, 7, 26, 2, 2, 233, 235, 5, 14, 8, 2, 234, 233, 3, 2, 2, 2, 235,
	238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 240,
	3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 241, 5, 16, 9, 2, 240, 239, 3, 2,
	2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 244, 5, 20, 11,
	2, 243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245,
	247, 7, 27, 2, 2, 246, 232, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 13,
	3, 2, 2, 2, 248, 253, 7, 34, 2, 2, 249, 250, 7, 26, 2, 2, 250, 251, 5,
	20, 11, 2, 251, 252, 7, 27, 2, 2, 252, 254, 3, 2, 2, 2, 253, 249, 3, 2,
	2, 2, 253, 254, 3, 2, 2, 2, 254, 270, 3, 2, 2, 2, 255, 260, 7, 35, 2, 2,
	256, 257, 7, 26, 2, 2, 257, 258, 5, 20, 11, 2, 258, 259, 7, 27, 2, 2, 259,
	261, 3, 2, 2, 2, 260, 256, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 270,
	3, 2, 2, 2, 262, 267, 7, 36, 2, 2, 263, 264, 7, 26, 2, 2, 264, 265, 5,
	20, 11, 2, 265, 266, 7, 27, 2, 2, 266, 268, 3, 2, 2, 2, 267, 263, 3, 2,
	2, 2, 267, 268, 3, 2, 2, 2, 268, 270, 3, 2, 2, 2, 269, 248, 3, 2, 2, 2,
	269, 255, 3, 2, 2, 2, 269, 262, 3, 2, 2, 2, 270, 15, 3, 2, 2, 2, 271, 282,
	7, 42, 2, 2, 272, 273, 7, 43, 2, 2, 273, 275, 7, 26, 2, 2, 274, 276, 5,
	18, 10, 2, 275, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275, 3, 2,
	2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 7, 27, 2, 2,
	280, 282, 3, 2, 2, 2, 281, 271, 3, 2, 2, 2, 281, 272, 3, 2, 2, 2, 282,
	17, 3, 2, 2, 2, 283, 292, 7, 44, 2, 2, 284, 286, 7, 26, 2, 2, 285, 287,
	5, 18, 10, 2, 286, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 286, 3,
	2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 7, 27, 2,
	2, 291, 293, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293,
	19, 3, 2, 2, 2, 294, 323, 7, 47, 2, 2, 295, 323, 7, 48, 2, 2, 296, 323,
	7, 49, 2, 2, 297, 323, 7, 50, 2, 2, 298, 323, 7, 51, 2, 2, 299, 308, 7,
	52, 2, 2, 300, 302, 7, 26, 2, 2, 301, 303, 5, 20, 11, 2, 302, 301, 3, 2,
	2, 2, 303, 304, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2,
	305, 306, 3, 2, 2, 2, 306, 307, 7, 27, 2, 2, 307, 309, 3, 2, 2, 2, 308,
	300, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 323, 3, 2, 2, 2, 310, 319,
	7, 53, 2, 2, 311, 313, 7, 26, 2, 2, 312, 314, 5, 20, 11, 2, 313, 312, 3,
	2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 315, 316, 3, 2, 2,
	2, 316, 317, 3, 2, 2, 2, 317, 318, 7, 27, 2, 2, 318, 320, 3, 2, 2, 2, 319,
	311, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 323, 3, 2, 2, 2, 321, 323,
	7, 29, 2, 2, 322, 294, 3, 2, 2, 2, 322, 295, 3, 2, 2, 2, 322, 296, 3, 2,
	2, 2, 322, 297, 3, 2, 2, 2, 322, 298, 3, 2, 2, 2, 322, 299, 3, 2, 2, 2,
	322, 310, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323, 21, 3, 2, 2, 2, 56, 40,
	46, 52, 56, 63, 66, 68, 73, 80, 84, 89, 93, 100, 104, 109, 113, 120, 124,
	127, 132, 136, 143, 147, 150, 155, 159, 178, 184, 191, 197, 203, 208, 213,
	217, 224, 227, 229, 236, 240, 243, 246, 253, 260, 267, 269, 277, 281, 288,
	292, 304, 308, 315, 319, 322,
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
	ERROR() antlr.TerminalNode
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

func (s *ModuleContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlWiERROR, 0)
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

	var _alt int

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiModule:
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
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(41)
						p.Import_()
					}

				}
				p.SetState(46)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
			}
			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiStruct-27))|(1<<(AdlWiUnion-27))|(1<<(AdlWiNewtype-27))|(1<<(AdlWiType-27))|(1<<(AdlWiModuleAnno-27))|(1<<(AdlWiDeclAnno-27))|(1<<(AdlWiFieldAnno-27)))) != 0 {
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

	case AdlWiERROR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(AdlWiERROR)
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(57)
				p.Match(AdlWiDOWN)
			}
			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_alt = 1 + 1
			for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1 + 1:
					p.SetState(58)
					p.MatchWildcard()

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(61)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
			}
			{
				p.SetState(63)
				p.Match(AdlWiUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type ImportErrorContext struct {
	*Import_Context
}

func NewImportErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportErrorContext {
	var p = new(ImportErrorContext)

	p.Import_Context = NewEmptyImport_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Import_Context))

	return p
}

type IImportErrorContext interface {
	//Current rule
	IImport_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ERROR() antlr.TerminalNode
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

func (*ImportErrorContext) IsImportErrorContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlWiERROR, 0)
}

func (s *ImportErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportErrorEntryListener); ok {
		listenerT.EnterImportError(s)
	}
}

func (s *ImportErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportErrorExitListener); ok {
		listenerT.ExitImportError(s)
	}
}

func (s *ImportErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportError != nil {
		h.ImportError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportErrorContextVisitor:
		return t.VisitImportError(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiImportModule:
		localctx = NewImportModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(AdlWiImportModule)
		}

	case AdlWiImportScopedName:
		localctx = NewImportScopedModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(AdlWiImportScopedName)
		}

	case AdlWiERROR:
		localctx = NewImportErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(AdlWiERROR)
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

type TLDErrorContext struct {
	*TldContext
}

func NewTLDErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TLDErrorContext {
	var p = new(TLDErrorContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type ITLDErrorContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ERROR() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
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

func (*TLDErrorContext) IsTLDErrorContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TLDErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TLDErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlWiERROR, 0)
}

func (s *TLDErrorContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWiDOWN, 0)
}

func (s *TLDErrorContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWiUP, 0)
}

func (s *TLDErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TLDErrorEntryListener); ok {
		listenerT.EnterTLDError(s)
	}
}

func (s *TLDErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TLDErrorExitListener); ok {
		listenerT.ExitTLDError(s)
	}
}

func (s *TLDErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TLDError != nil {
		h.TLDError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TLDErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TLDErrorContextVisitor:
		return t.VisitTLDError(s, delegate, args...)
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

	var _alt int

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			var _m = p.Match(AdlWiStruct)
			localctx.(*StructContext).tok = _m

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(74)
				p.Match(AdlWiDOWN)
			}
			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(75)
					p.Annotation()
				}

				p.SetState(80)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(81)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWiField {
				{
					p.SetState(84)
					p.NameBody()
				}

				p.SetState(89)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(90)
				p.Match(AdlWiUP)
			}

		}

	case 2:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			var _m = p.Match(AdlWiUnion)
			localctx.(*UnionContext).tok = _m

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(94)
				p.Match(AdlWiDOWN)
			}
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(95)
					p.Annotation()
				}

				p.SetState(100)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(101)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWiField {
				{
					p.SetState(104)
					p.NameBody()
				}

				p.SetState(109)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(110)
				p.Match(AdlWiUP)
			}

		}

	case 3:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			var _m = p.Match(AdlWiType)
			localctx.(*TypeContext).tok = _m

		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(114)
				p.Match(AdlWiDOWN)
			}
			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(115)
					p.Annotation()
				}

				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(121)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
				{
					p.SetState(124)
					p.TypeExpr_()
				}

			}
			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiJsonStr-27))|(1<<(AdlWiJsonBool-27))|(1<<(AdlWiJsonNull-27))|(1<<(AdlWiJsonInt-27))|(1<<(AdlWiJsonFloat-27))|(1<<(AdlWiJsonArray-27))|(1<<(AdlWiJsonObj-27)))) != 0 {
				{
					p.SetState(127)
					p.JsonVal()
				}

				p.SetState(132)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(133)
				p.Match(AdlWiUP)
			}

		}

	case 4:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			var _m = p.Match(AdlWiNewtype)
			localctx.(*NewtypeContext).tok = _m

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(137)
				p.Match(AdlWiDOWN)
			}
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
				{
					p.SetState(138)
					p.Annotation()
				}

				p.SetState(143)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeParam {
				{
					p.SetState(144)
					p.Match(AdlWiTypeParam)
				}

			}
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
				{
					p.SetState(147)
					p.TypeExpr_()
				}

			}
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiJsonStr-27))|(1<<(AdlWiJsonBool-27))|(1<<(AdlWiJsonNull-27))|(1<<(AdlWiJsonInt-27))|(1<<(AdlWiJsonFloat-27))|(1<<(AdlWiJsonArray-27))|(1<<(AdlWiJsonObj-27)))) != 0 {
				{
					p.SetState(150)
					p.JsonVal()
				}

				p.SetState(155)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(156)
				p.Match(AdlWiUP)
			}

		}

	case 5:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(159)
			var _m = p.Match(AdlWiModuleAnno)
			localctx.(*ModAnnoContext).tok = _m

		}
		{
			p.SetState(160)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(161)
			p.JsonVal()
		}
		{
			p.SetState(162)
			p.Match(AdlWiUP)
		}

	case 6:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(164)
			var _m = p.Match(AdlWiDeclAnno)
			localctx.(*DeclAnnoContext).tok = _m

		}
		{
			p.SetState(165)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(166)
			p.JsonVal()
		}
		{
			p.SetState(167)
			p.Match(AdlWiUP)
		}

	case 7:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(169)
			var _m = p.Match(AdlWiFieldAnno)
			localctx.(*FieldAnnoContext).tok = _m

		}
		{
			p.SetState(170)
			p.Match(AdlWiDOWN)
		}
		{
			p.SetState(171)
			p.JsonVal()
		}
		{
			p.SetState(172)
			p.Match(AdlWiUP)
		}

	case 8:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		p.SetState(176)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AdlWiStruct:
			{
				p.SetState(174)
				var _m = p.Match(AdlWiStruct)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		case AdlWiUnion:
			{
				p.SetState(175)
				var _m = p.Match(AdlWiUnion)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(178)
			p.Match(AdlWiDOWN)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(179)
				p.Annotation()
			}

			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(185)
			p.Match(AdlWiERROR)
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlWiField {
			{
				p.SetState(186)
				p.NameBody()
			}

			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(192)
			p.Match(AdlWiUP)
		}

	case 9:
		localctx = NewTypeParamErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		p.SetState(195)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case AdlWiType:
			{
				p.SetState(193)
				var _m = p.Match(AdlWiType)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		case AdlWiNewtype:
			{
				p.SetState(194)
				var _m = p.Match(AdlWiNewtype)
				localctx.(*TypeParamErrorContext).tok = _m

			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(197)
			p.Match(AdlWiDOWN)
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(198)
				p.Annotation()
			}

			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(204)
			p.Match(AdlWiERROR)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
			{
				p.SetState(205)
				p.TypeExpr_()
			}

		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(208)
					p.JsonVal()
				}

			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(214)
				p.Match(AdlWiUP)
			}

		}

	case 10:
		localctx = NewTLDErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(217)
			p.Match(AdlWiERROR)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(218)
				p.Match(AdlWiDOWN)
			}
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_alt = 1 + 1
			for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1 + 1:
					p.SetState(219)
					p.MatchWildcard()

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(222)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
			}
			{
				p.SetState(224)
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
		p.SetState(229)
		var _m = p.Match(AdlWiField)
		localctx.(*FieldContext).tok = _m

	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWiDOWN {
		{
			p.SetState(230)
			p.Match(AdlWiDOWN)
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWiAnnotation-32))|(1<<(AdlWiAnnotationNotScoped-32))|(1<<(AdlWiAnnotationScoped-32)))) != 0 {
			{
				p.SetState(231)
				p.Annotation()
			}

			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiTypeExprPrimOrParam || _la == AdlWiTypeExprTypeExpr {
			{
				p.SetState(237)
				p.TypeExpr_()
			}

		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiJsonStr-27))|(1<<(AdlWiJsonBool-27))|(1<<(AdlWiJsonNull-27))|(1<<(AdlWiJsonInt-27))|(1<<(AdlWiJsonFloat-27))|(1<<(AdlWiJsonArray-27))|(1<<(AdlWiJsonObj-27)))) != 0 {
			{
				p.SetState(240)
				p.JsonVal()
			}

		}
		{
			p.SetState(243)
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

	p.SetState(267)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiAnnotation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			var _m = p.Match(AdlWiAnnotation)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(247)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(248)
				p.JsonVal()
			}
			{
				p.SetState(249)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiAnnotationNotScoped:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			var _m = p.Match(AdlWiAnnotationNotScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(254)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(255)
				p.JsonVal()
			}
			{
				p.SetState(256)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiAnnotationScoped:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			var _m = p.Match(AdlWiAnnotationScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(261)
				p.Match(AdlWiDOWN)
			}
			{
				p.SetState(262)
				p.JsonVal()
			}
			{
				p.SetState(263)
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

	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiTypeExprPrimOrParam:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(269)
			var _m = p.Match(AdlWiTypeExprPrimOrParam)
			localctx.(*TypeExpr_Context).tok = _m

		}

	case AdlWiTypeExprTypeExpr:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			var _m = p.Match(AdlWiTypeExprTypeExpr)
			localctx.(*TypeExpr_Context).tok = _m

		}

		{
			p.SetState(271)
			p.Match(AdlWiDOWN)
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWiTypeExprElem {
			{
				p.SetState(272)
				p.TypeExprElem_()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(277)
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
		p.SetState(281)
		var _m = p.Match(AdlWiTypeExprElem)
		localctx.(*TypeParamsContext).tok = _m

	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWiDOWN {
		{
			p.SetState(282)
			p.Match(AdlWiDOWN)
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWiTypeExprElem {
			{
				p.SetState(283)
				p.TypeExprElem_()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(288)
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

type JsonErrorContext struct {
	*JsonValContext
}

func NewJsonErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonErrorContext {
	var p = new(JsonErrorContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonErrorContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ERROR() antlr.TerminalNode
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

func (*JsonErrorContext) IsJsonErrorContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlWiERROR, 0)
}

func (s *JsonErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonErrorEntryListener); ok {
		listenerT.EnterJsonError(s)
	}
}

func (s *JsonErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonErrorExitListener); ok {
		listenerT.ExitJsonError(s)
	}
}

func (s *JsonErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWiHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonError != nil {
		h.JsonError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonErrorContextVisitor:
		return t.VisitJsonError(s, delegate, args...)
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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWiJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			var _m = p.Match(AdlWiJsonStr)
			localctx.(*JsonStrContext).tok = _m

		}

	case AdlWiJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			var _m = p.Match(AdlWiJsonBool)
			localctx.(*JsonBoolContext).tok = _m

		}

	case AdlWiJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(294)
			var _m = p.Match(AdlWiJsonNull)
			localctx.(*JsonNullContext).tok = _m

		}

	case AdlWiJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(295)
			var _m = p.Match(AdlWiJsonInt)
			localctx.(*JsonIntContext).tok = _m

		}

	case AdlWiJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(296)
			var _m = p.Match(AdlWiJsonFloat)
			localctx.(*JsonFloatContext).tok = _m

		}

	case AdlWiJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(297)
			var _m = p.Match(AdlWiJsonArray)
			localctx.(*JsonArrayContext).tok = _m

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(298)
				p.Match(AdlWiDOWN)
			}
			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiJsonStr-27))|(1<<(AdlWiJsonBool-27))|(1<<(AdlWiJsonNull-27))|(1<<(AdlWiJsonInt-27))|(1<<(AdlWiJsonFloat-27))|(1<<(AdlWiJsonArray-27))|(1<<(AdlWiJsonObj-27)))) != 0) {
				{
					p.SetState(299)
					p.JsonVal()
				}

				p.SetState(302)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(304)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(308)
			var _m = p.Match(AdlWiJsonObj)
			localctx.(*JsonObjContext).tok = _m

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWiDOWN {
			{
				p.SetState(309)
				p.Match(AdlWiDOWN)
			}
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlWiERROR-27))|(1<<(AdlWiJsonStr-27))|(1<<(AdlWiJsonBool-27))|(1<<(AdlWiJsonNull-27))|(1<<(AdlWiJsonInt-27))|(1<<(AdlWiJsonFloat-27))|(1<<(AdlWiJsonArray-27))|(1<<(AdlWiJsonObj-27)))) != 0) {
				{
					p.SetState(310)
					p.JsonVal()
				}

				p.SetState(313)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(315)
				p.Match(AdlWiUP)
			}

		}

	case AdlWiERROR:
		localctx = NewJsonErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(319)
			p.Match(AdlWiERROR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
