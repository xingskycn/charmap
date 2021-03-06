
package charmap

func init() {

	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // NULL
		'\x01':	'\u0001',	 // START OF HEADING
		'\x02':	'\u0002',	 // START OF TEXT
		'\x03':	'\u0003',	 // END OF TEXT
		'\x04':	'\u0004',	 // END OF TRANSMISSION
		'\x05':	'\u0005',	 // ENQUIRY
		'\x06':	'\u0006',	 // ACKNOWLEDGE
		'\x07':	'\u0007',	 // BELL
		'\x08':	'\u0008',	 // BACKSPACE
		'\x09':	'\u0009',	 // HORIZONTAL TABULATION
		'\x0A':	'\u000A',	 // LINE FEED
		'\x0B':	'\u000B',	 // VERTICAL TABULATION
		'\x0C':	'\u000C',	 // FORM FEED
		'\x0D':	'\u000D',	 // CARRIAGE RETURN
		'\x0E':	'\u000E',	 // SHIFT OUT
		'\x0F':	'\u000F',	 // SHIFT IN
		'\x10':	'\u0010',	 // DATA LINK ESCAPE
		'\x11':	'\u0011',	 // DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // DEVICE CONTROL THREE
		'\x14':	'\u0014',	 // DEVICE CONTROL FOUR
		'\x15':	'\u0015',	 // NEGATIVE ACKNOWLEDGE
		'\x16':	'\u0016',	 // SYNCHRONOUS IDLE
		'\x17':	'\u0017',	 // END OF TRANSMISSION BLOCK
		'\x18':	'\u0018',	 // CANCEL
		'\x19':	'\u0019',	 // END OF MEDIUM
		'\x1A':	'\u001A',	 // SUBSTITUTE
		'\x1B':	'\u001B',	 // ESCAPE
		'\x1C':	'\u001C',	 // FILE SEPARATOR
		'\x1D':	'\u001D',	 // GROUP SEPARATOR
		'\x1E':	'\u001E',	 // RECORD SEPARATOR
		'\x1F':	'\u001F',	 // UNIT SEPARATOR
		'\x20':	'\u0020',	 // SPACE
		'\x21':	'\u0021',	 // EXCLAMATION MARK
		'\x22':	'\u0022',	 // QUOTATION MARK
		'\x23':	'\u0023',	 // NUMBER SIGN
		'\x24':	'\u0024',	 // DOLLAR SIGN
		'\x25':	'\u0025',	 // PERCENT SIGN
		'\x26':	'\u0026',	 // AMPERSAND
		'\x27':	'\u0027',	 // APOSTROPHE
		'\x28':	'\u0028',	 // LEFT PARENTHESIS
		'\x29':	'\u0029',	 // RIGHT PARENTHESIS
		'\x2A':	'\u002A',	 // ASTERISK
		'\x2B':	'\u002B',	 // PLUS SIGN
		'\x2C':	'\u002C',	 // COMMA
		'\x2D':	'\u002D',	 // HYPHEN-MINUS
		'\x2E':	'\u002E',	 // FULL STOP
		'\x2F':	'\u002F',	 // SOLIDUS
		'\x30':	'\u0030',	 // DIGIT ZERO
		'\x31':	'\u0031',	 // DIGIT ONE
		'\x32':	'\u0032',	 // DIGIT TWO
		'\x33':	'\u0033',	 // DIGIT THREE
		'\x34':	'\u0034',	 // DIGIT FOUR
		'\x35':	'\u0035',	 // DIGIT FIVE
		'\x36':	'\u0036',	 // DIGIT SIX
		'\x37':	'\u0037',	 // DIGIT SEVEN
		'\x38':	'\u0038',	 // DIGIT EIGHT
		'\x39':	'\u0039',	 // DIGIT NINE
		'\x3A':	'\u003A',	 // COLON
		'\x3B':	'\u003B',	 // SEMICOLON
		'\x3C':	'\u003C',	 // LESS-THAN SIGN
		'\x3D':	'\u003D',	 // EQUALS SIGN
		'\x3E':	'\u003E',	 // GREATER-THAN SIGN
		'\x3F':	'\u003F',	 // QUESTION MARK
		'\x40':	'\u0040',	 // COMMERCIAL AT
		'\x41':	'\u0041',	 // LATIN CAPITAL LETTER A
		'\x42':	'\u0042',	 // LATIN CAPITAL LETTER B
		'\x43':	'\u0043',	 // LATIN CAPITAL LETTER C
		'\x44':	'\u0044',	 // LATIN CAPITAL LETTER D
		'\x45':	'\u0045',	 // LATIN CAPITAL LETTER E
		'\x46':	'\u0046',	 // LATIN CAPITAL LETTER F
		'\x47':	'\u0047',	 // LATIN CAPITAL LETTER G
		'\x48':	'\u0048',	 // LATIN CAPITAL LETTER H
		'\x49':	'\u0049',	 // LATIN CAPITAL LETTER I
		'\x4A':	'\u004A',	 // LATIN CAPITAL LETTER J
		'\x4B':	'\u004B',	 // LATIN CAPITAL LETTER K
		'\x4C':	'\u004C',	 // LATIN CAPITAL LETTER L
		'\x4D':	'\u004D',	 // LATIN CAPITAL LETTER M
		'\x4E':	'\u004E',	 // LATIN CAPITAL LETTER N
		'\x4F':	'\u004F',	 // LATIN CAPITAL LETTER O
		'\x50':	'\u0050',	 // LATIN CAPITAL LETTER P
		'\x51':	'\u0051',	 // LATIN CAPITAL LETTER Q
		'\x52':	'\u0052',	 // LATIN CAPITAL LETTER R
		'\x53':	'\u0053',	 // LATIN CAPITAL LETTER S
		'\x54':	'\u0054',	 // LATIN CAPITAL LETTER T
		'\x55':	'\u0055',	 // LATIN CAPITAL LETTER U
		'\x56':	'\u0056',	 // LATIN CAPITAL LETTER V
		'\x57':	'\u0057',	 // LATIN CAPITAL LETTER W
		'\x58':	'\u0058',	 // LATIN CAPITAL LETTER X
		'\x59':	'\u0059',	 // LATIN CAPITAL LETTER Y
		'\x5A':	'\u005A',	 // LATIN CAPITAL LETTER Z
		'\x5B':	'\u005B',	 // LEFT SQUARE BRACKET
		'\x5C':	'\u005C',	 // REVERSE SOLIDUS
		'\x5D':	'\u005D',	 // RIGHT SQUARE BRACKET
		'\x5E':	'\u005E',	 // CIRCUMFLEX ACCENT
		'\x5F':	'\u005F',	 // LOW LINE
		'\x60':	'\u0060',	 // GRAVE ACCENT
		'\x61':	'\u0061',	 // LATIN SMALL LETTER A
		'\x62':	'\u0062',	 // LATIN SMALL LETTER B
		'\x63':	'\u0063',	 // LATIN SMALL LETTER C
		'\x64':	'\u0064',	 // LATIN SMALL LETTER D
		'\x65':	'\u0065',	 // LATIN SMALL LETTER E
		'\x66':	'\u0066',	 // LATIN SMALL LETTER F
		'\x67':	'\u0067',	 // LATIN SMALL LETTER G
		'\x68':	'\u0068',	 // LATIN SMALL LETTER H
		'\x69':	'\u0069',	 // LATIN SMALL LETTER I
		'\x6A':	'\u006A',	 // LATIN SMALL LETTER J
		'\x6B':	'\u006B',	 // LATIN SMALL LETTER K
		'\x6C':	'\u006C',	 // LATIN SMALL LETTER L
		'\x6D':	'\u006D',	 // LATIN SMALL LETTER M
		'\x6E':	'\u006E',	 // LATIN SMALL LETTER N
		'\x6F':	'\u006F',	 // LATIN SMALL LETTER O
		'\x70':	'\u0070',	 // LATIN SMALL LETTER P
		'\x71':	'\u0071',	 // LATIN SMALL LETTER Q
		'\x72':	'\u0072',	 // LATIN SMALL LETTER R
		'\x73':	'\u0073',	 // LATIN SMALL LETTER S
		'\x74':	'\u0074',	 // LATIN SMALL LETTER T
		'\x75':	'\u0075',	 // LATIN SMALL LETTER U
		'\x76':	'\u0076',	 // LATIN SMALL LETTER V
		'\x77':	'\u0077',	 // LATIN SMALL LETTER W
		'\x78':	'\u0078',	 // LATIN SMALL LETTER X
		'\x79':	'\u0079',	 // LATIN SMALL LETTER Y
		'\x7A':	'\u007A',	 // LATIN SMALL LETTER Z
		'\x7B':	'\u007B',	 // LEFT CURLY BRACKET
		'\x7C':	'\u007C',	 // VERTICAL LINE
		'\x7D':	'\u007D',	 // RIGHT CURLY BRACKET
		'\x7E':	'\u007E',	 // TILDE
		'\x7F':	'\u007F',	 // DELETE
		'\x80':	'\u05D0',	 // HEBREW LETTER ALEF
		'\x81':	'\u05D1',	 // HEBREW LETTER BET
		'\x82':	'\u05D2',	 // HEBREW LETTER GIMEL
		'\x83':	'\u05D3',	 // HEBREW LETTER DALET
		'\x84':	'\u05D4',	 // HEBREW LETTER HE
		'\x85':	'\u05D5',	 // HEBREW LETTER VAV
		'\x86':	'\u05D6',	 // HEBREW LETTER ZAYIN
		'\x87':	'\u05D7',	 // HEBREW LETTER HET
		'\x88':	'\u05D8',	 // HEBREW LETTER TET
		'\x89':	'\u05D9',	 // HEBREW LETTER YOD
		'\x8A':	'\u05DA',	 // HEBREW LETTER FINAL KAF
		'\x8B':	'\u05DB',	 // HEBREW LETTER KAF
		'\x8C':	'\u05DC',	 // HEBREW LETTER LAMED
		'\x8D':	'\u05DD',	 // HEBREW LETTER FINAL MEM
		'\x8E':	'\u05DE',	 // HEBREW LETTER MEM
		'\x8F':	'\u05DF',	 // HEBREW LETTER FINAL NUN
		'\x90':	'\u05E0',	 // HEBREW LETTER NUN
		'\x91':	'\u05E1',	 // HEBREW LETTER SAMEKH
		'\x92':	'\u05E2',	 // HEBREW LETTER AYIN
		'\x93':	'\u05E3',	 // HEBREW LETTER FINAL PE
		'\x94':	'\u05E4',	 // HEBREW LETTER PE
		'\x95':	'\u05E5',	 // HEBREW LETTER FINAL TSADI
		'\x96':	'\u05E6',	 // HEBREW LETTER TSADI
		'\x97':	'\u05E7',	 // HEBREW LETTER QOF
		'\x98':	'\u05E8',	 // HEBREW LETTER RESH
		'\x99':	'\u05E9',	 // HEBREW LETTER SHIN
		'\x9A':	'\u05EA',	 // HEBREW LETTER TAV
		// '\x9B' UNDEFINED
		'\x9C':	'\u00A3',	 // POUND SIGN
		// '\x9D' UNDEFINED
		'\x9E':	'\u00D7',	 // MULTIPLICATION SIGN
		// '\x9F' UNDEFINED
		// '\xA0' UNDEFINED
		// '\xA1' UNDEFINED
		// '\xA2' UNDEFINED
		// '\xA3' UNDEFINED
		// '\xA4' UNDEFINED
		// '\xA5' UNDEFINED
		// '\xA6' UNDEFINED
		// '\xA7' UNDEFINED
		// '\xA8' UNDEFINED
		'\xA9':	'\u00AE',	 // REGISTERED SIGN
		'\xAA':	'\u00AC',	 // NOT SIGN
		'\xAB':	'\u00BD',	 // VULGAR FRACTION ONE HALF
		'\xAC':	'\u00BC',	 // VULGAR FRACTION ONE QUARTER
		// '\xAD' UNDEFINED
		'\xAE':	'\u00AB',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xAF':	'\u00BB',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xB0':	'\u2591',	 // LIGHT SHADE
		'\xB1':	'\u2592',	 // MEDIUM SHADE
		'\xB2':	'\u2593',	 // DARK SHADE
		'\xB3':	'\u2502',	 // BOX DRAWINGS LIGHT VERTICAL
		'\xB4':	'\u2524',	 // BOX DRAWINGS LIGHT VERTICAL AND LEFT
		// '\xB5' UNDEFINED
		// '\xB6' UNDEFINED
		// '\xB7' UNDEFINED
		'\xB8':	'\u00A9',	 // COPYRIGHT SIGN
		'\xB9':	'\u2563',	 // BOX DRAWINGS DOUBLE VERTICAL AND LEFT
		'\xBA':	'\u2551',	 // BOX DRAWINGS DOUBLE VERTICAL
		'\xBB':	'\u2557',	 // BOX DRAWINGS DOUBLE DOWN AND LEFT
		'\xBC':	'\u255D',	 // BOX DRAWINGS DOUBLE UP AND LEFT
		'\xBD':	'\u00A2',	 // CENT SIGN
		'\xBE':	'\u00A5',	 // YEN SIGN
		'\xBF':	'\u2510',	 // BOX DRAWINGS LIGHT DOWN AND LEFT
		'\xC0':	'\u2514',	 // BOX DRAWINGS LIGHT UP AND RIGHT
		'\xC1':	'\u2534',	 // BOX DRAWINGS LIGHT UP AND HORIZONTAL
		'\xC2':	'\u252C',	 // BOX DRAWINGS LIGHT DOWN AND HORIZONTAL
		'\xC3':	'\u251C',	 // BOX DRAWINGS LIGHT VERTICAL AND RIGHT
		'\xC4':	'\u2500',	 // BOX DRAWINGS LIGHT HORIZONTAL
		'\xC5':	'\u253C',	 // BOX DRAWINGS LIGHT VERTICAL AND HORIZONTAL
		// '\xC6' UNDEFINED
		// '\xC7' UNDEFINED
		'\xC8':	'\u255A',	 // BOX DRAWINGS DOUBLE UP AND RIGHT
		'\xC9':	'\u2554',	 // BOX DRAWINGS DOUBLE DOWN AND RIGHT
		'\xCA':	'\u2569',	 // BOX DRAWINGS DOUBLE UP AND HORIZONTAL
		'\xCB':	'\u2566',	 // BOX DRAWINGS DOUBLE DOWN AND HORIZONTAL
		'\xCC':	'\u2560',	 // BOX DRAWINGS DOUBLE VERTICAL AND RIGHT
		'\xCD':	'\u2550',	 // BOX DRAWINGS DOUBLE HORIZONTAL
		'\xCE':	'\u256C',	 // BOX DRAWINGS DOUBLE VERTICAL AND HORIZONTAL
		'\xCF':	'\u00A4',	 // CURRENCY SIGN
		// '\xD0' UNDEFINED
		// '\xD1' UNDEFINED
		// '\xD2' UNDEFINED
		// '\xD3' UNDEFINEDS
		// '\xD4' UNDEFINED
		// '\xD5' UNDEFINED
		// '\xD6' UNDEFINEDE
		// '\xD7' UNDEFINED
		// '\xD8' UNDEFINED
		'\xD9':	'\u2518',	 // BOX DRAWINGS LIGHT UP AND LEFT
		'\xDA':	'\u250C',	 // BOX DRAWINGS LIGHT DOWN AND RIGHT
		'\xDB':	'\u2588',	 // FULL BLOCK
		'\xDC':	'\u2584',	 // LOWER HALF BLOCK
		'\xDD':	'\u00A6',	 // BROKEN BAR
		// '\xDE' UNDEFINED
		'\xDF':	'\u2580',	 // UPPER HALF BLOCK
		// '\xE0' UNDEFINED
		// '\xE1' UNDEFINED
		// '\xE2' UNDEFINED
		// '\xE3' UNDEFINED
		// '\xE4' UNDEFINED
		// '\xE5' UNDEFINED
		'\xE6':	'\u00B5',	 // MICRO SIGN
		// '\xE7' UNDEFINED
		// '\xE8' UNDEFINED
		// '\xE9' UNDEFINED
		// '\xEA' UNDEFINED
		// '\xEB' UNDEFINED
		// '\xEC' UNDEFINED
		// '\xED' UNDEFINED
		'\xEE':	'\u00AF',	 // MACRON
		'\xEF':	'\u00B4',	 // ACUTE ACCENT
		'\xF0':	'\u00AD',	 // SOFT HYPHEN
		'\xF1':	'\u00B1',	 // PLUS-MINUS SIGN
		'\xF2':	'\u2017',	 // DOUBLE LOW LINE
		'\xF3':	'\u00BE',	 // VULGAR FRACTION THREE QUARTERS
		'\xF4':	'\u00B6',	 // PILCROW SIGN
		'\xF5':	'\u00A7',	 // SECTION SIGN
		'\xF6':	'\u00F7',	 // DIVISION SIGN
		'\xF7':	'\u00B8',	 // CEDILLA
		'\xF8':	'\u00B0',	 // DEGREE SIGN
		'\xF9':	'\u00A8',	 // DIAERESIS
		'\xFA':	'\u00B7',	 // MIDDLE DOT
		'\xFB':	'\u00B9',	 // SUPERSCRIPT ONE
		'\xFC':	'\u00B3',	 // SUPERSCRIPT THREE
		'\xFD':	'\u00B2',	 // SUPERSCRIPT TWO
		'\xFE':	'\u25A0',	 // BLACK SQUARE
		'\xFF':	'\u00A0',	 // NO-BREAK SPACE

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	newCodec := &codecMap8Bit{EncodeMap: charmapEncode, DecodeMap: charmapDecode}

	register(newCodec, "CP856", "CP-856", "856")

}
