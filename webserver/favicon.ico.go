// return a default favicon in case user html files don't point to one

package webserver

func faviconData() string {
    return "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A\x00\x00\x00\x0D\x49\x48\x44\x52" +
           "\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1F\xF3\xFF" +
           "\x61\x00\x00\x01\x3D\x49\x44\x41\x54\x38\x4F\x63\x0C\xFD\xF1\x5A" +
           "\x9D\x81\x02\xC0\x18\xF8\xE4\xB5\x1A\x05\xFA\x19\x18\xFD\xAE\xBC" +
           "\x52\xA1\xC8\x00\xCF\xA3\x1F\x94\xF1\x19\xB0\xDD\x5A\xE0\x0E\x50" +
           "\x0D\xDC\x92\x3D\x3E\x2A\x8D\xBF\xDF\xBF\x89\x06\xE9\x01\x89\x33" +
           "\xBA\x6C\xF9\xA0\x84\xCF\x80\x3D\x3E\x02\x77\x81\x6A\xE0\x96\xA0" +
           "\xF3\x19\xED\x16\x7D\x50\xC4\x67\xC0\xA1\x38\x81\x7B\x40\x35\x70" +
           "\x4B\xD0\xF9\x8C\x56\x53\xDE\xCB\x83\x0C\x38\x96\x23\xF8\x80\x95" +
           "\x4F\x72\xE2\xEF\x4F\xCF\xF3\x41\x7C\xA0\xB8\x02\x4C\x1C\x99\x8D" +
           "\x6C\x19\x48\x9C\xD1\xB8\xE3\xBB\xDC\xD9\x0A\xCE\x87\x40\x1A\x6C" +
           "\x10\x08\x20\xF3\xF1\xC9\x81\xD4\x32\xEA\x14\x7D\x93\xBD\xD2\xC7" +
           "\xF5\x88\x81\x89\xED\xAA\x4E\xC1\x07\x4F\x74\xEF\x80\xE4\x80\x6A" +
           "\xE4\x60\xE2\xE8\x7C\x46\x95\x94\x6F\x32\x20\xC9\x3B\x73\xB8\x1E" +
           "\xC3\x14\x01\xC5\x64\x61\x6C\x90\x38\x3E\x3E\xA3\x6C\xC4\x57\x29" +
           "\x74\x5B\x1F\xAF\xE0\x7E\x0A\x14\x97\x06\x89\x23\xB3\xB1\xF1\x19" +
           "\x45\xDC\xBF\x4A\xBE\xD9\xC9\xFD\x0C\x48\xC3\x0D\x7A\xB3\x93\xF7" +
           "\xA2\x88\xFB\x67\x7D\x90\x06\x4C\x39\x54\xB5\x8C\xA2\xA6\xFF\x25" +
           "\x5E\x9F\x66\x3D\xC4\xC0\xF0\x47\x15\xC8\x96\x04\x69\x7A\x7D\x9A" +
           "\xF1\x39\x36\x36\xBA\x1C\x38\x10\xB9\x15\xFF\x8B\xC3\xBC\xF0\xF5" +
           "\x3E\xE3\x0B\x06\x06\xD6\x2B\xDC\x8A\xBF\x5C\xD0\xBD\x85\x8B\xCF" +
           "\xC8\x23\xF1\x5F\x94\x58\xC5\xD8\xD4\x31\x32\x30\xFC\x17\xA1\xC4" +
           "\x00\x00\xEF\xED\xA1\xD4\x4D\x65\x4F\x3B\x00\x00\x00\x00\x49\x45" +
           "\x4E\x44\xAE\x42\x60\x82"
}


//  0000: 89 50 4E 47 0D 0A 1A 0A 00 00 00 0D 49 48 44 52 	.PNG........IHDR
//  0010: 00 00 00 10 00 00 00 10 08 06 00 00 00 1F F3 FF 	................
//  0020: 61 00 00 01 3D 49 44 41 54 38 4F 63 0C FD F1 5A 	a...=IDAT8Oc...Z
//  0030: 9D 81 02 C0 18 F8 E4 B5 1A 05 FA 19 18 FD AE BC 	................
//  0040: 52 A1 C8 00 CF A3 1F 94 F1 19 B0 DD 5A E0 0E 50 	R...........Z..P
//  0050: 0D DC 92 3D 3E 2A 8D BF DF BF 89 06 E9 01 89 33 	...=>*.........3
//  0060: BA 6C F9 A0 84 CF 80 3D 3E 02 77 81 6A E0 96 A0 	.l.....=>.w.j...
//  0070: F3 19 ED 16 7D 50 C4 67 C0 A1 38 81 7B 40 35 70 	....}P.g..8.{@5p
//  0080: 4B D0 F9 8C 56 53 DE CB 83 0C 38 96 23 F8 80 95 	K...VS....8.#...
//  0090: 4F 72 E2 EF 4F CF F3 41 7C A0 B8 02 4C 1C 99 8D 	Or..O..A|...L...
//  00A0: 6C 19 48 9C D1 B8 E3 BB DC D9 0A CE 87 40 1A 6C 	l.H..........@.l
//  00B0: 10 08 20 F3 F1 C9 81 D4 32 EA 14 7D 93 BD D2 C7 	.. .....2..}....
//  00C0: F5 88 81 89 ED AA 4E C1 07 4F 74 EF 80 E4 80 6A 	......N..Ot....j
//  00D0: E4 60 E2 E8 7C 46 95 94 6F 32 20 C9 3B 73 B8 1E 	.`..|F..o2 .;s..
//  00E0: C3 14 01 C5 64 61 6C 90 38 3E 3E A3 6C C4 57 29 	....dal.8>>.l.W)
//  00F0: 74 5B 1F AF E0 7E 0A 14 97 06 89 23 B3 B1 F1 19 	t[...~.....#....
//
//  0100: 45 DC BF 4A BE D9 C9 FD 0C 48 C3 0D 7A B3 93 F7 	E..J.....H..z...
//  0110: A2 88 FB 67 7D 90 06 4C 39 54 B5 8C A2 A6 FF 25 	...g}..L9T.....%
//  0120: 5E 9F 66 3D C4 C0 F0 47 15 C8 96 04 69 7A 7D 9A 	^.f=...G....iz}.
//  0130: F1 39 36 36 BA 1C 38 10 B9 15 FF 8B C3 BC F0 F5 	.966..8.........
//  0140: 3E E3 0B 06 06 D6 2B DC 8A BF 5C D0 BD 85 8B CF 	>.....+...\.....
//  0150: C8 23 F1 5F 94 58 C5 D8 D4 31 32 30 FC 17 A1 C4 	.#._.X...120....
//  0160: 00 00 EF ED A1 D4 4D 65 4F 3B 00 00 00 00 49 45 	......MeO;....IE
//  0170: 4E 44 AE 42 60 82                               	ND.B`.
