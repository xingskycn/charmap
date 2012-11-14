
package charmap

func init() {

	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // NULL
		'\x01':	'\u0001',	 // START OF HEADING
		'\x02':	'\u0002',	 // START OF TEXT
		'\x03':	'\u0003',	 // END OF TEXT
		'\x04':	'\u009C',	 // SELECT
		'\x05':	'\u0009',	 // HORIZONTAL TABULATION
		'\x06':	'\u0086',	 // REQUIRED NEW LINE
		'\x07':	'\u007F',	 // DELETE
		'\x08':	'\u0097',	 // GRAPHIC ESCAPE
		'\x09':	'\u008D',	 // SUPERSCRIPT
		'\x0A':	'\u008E',	 // REPEAT
		'\x0B':	'\u000B',	 // VERTICAL TABULATION
		'\x0C':	'\u000C',	 // FORM FEED
		'\x0D':	'\u000D',	 // CARRIAGE RETURN
		'\x0E':	'\u000E',	 // SHIFT OUT
		'\x0F':	'\u000F',	 // SHIFT IN
		'\x10':	'\u0010',	 // DATA LINK ESCAPE
		'\x11':	'\u0011',	 // DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // DEVICE CONTROL THREE
		'\x14':	'\u009D',	 // RESTORE/ENABLE PRESENTATION
		'\x15':	'\u0085',	 // NEW LINE
		'\x16':	'\u0008',	 // BACKSPACE
		'\x17':	'\u0087',	 // PROGRAM OPERATOR COMMUNICATION
		'\x18':	'\u0018',	 // CANCEL
		'\x19':	'\u0019',	 // END OF MEDIUM
		'\x1A':	'\u0092',	 // UNIT BACK SPACE
		'\x1B':	'\u008F',	 // CUSTOMER USE ONE
		'\x1C':	'\u001C',	 // FILE SEPARATOR
		'\x1D':	'\u001D',	 // GROUP SEPARATOR
		'\x1E':	'\u001E',	 // RECORD SEPARATOR
		'\x1F':	'\u001F',	 // UNIT SEPARATOR
		'\x20':	'\u0080',	 // DIGIT SELECT
		'\x21':	'\u0081',	 // START OF SIGNIFICANCE
		'\x22':	'\u0082',	 // FIELD SEPARATOR
		'\x23':	'\u0083',	 // WORD UNDERSCORE
		'\x24':	'\u0084',	 // BYPASS OR INHIBIT PRESENTATION
		'\x25':	'\u000A',	 // LINE FEED
		'\x26':	'\u0017',	 // END OF TRANSMISSION BLOCK
		'\x27':	'\u001B',	 // ESCAPE
		'\x28':	'\u0088',	 // SET ATTRIBUTE
		'\x29':	'\u0089',	 // START FIELD EXTENDED
		'\x2A':	'\u008A',	 // SET MODE OR SWITCH
		'\x2B':	'\u008B',	 // CONTROL SEQUENCE PREFIX
		'\x2C':	'\u008C',	 // MODIFY FIELD ATTRIBUTE
		'\x2D':	'\u0005',	 // ENQUIRY
		'\x2E':	'\u0006',	 // ACKNOWLEDGE
		'\x2F':	'\u0007',	 // BELL
		'\x30':	'\u0090',	 // <reserved>
		'\x31':	'\u0091',	 // <reserved>
		'\x32':	'\u0016',	 // SYNCHRONOUS IDLE
		'\x33':	'\u0093',	 // INDEX RETURN
		'\x34':	'\u0094',	 // PRESENTATION POSITION
		'\x35':	'\u0095',	 // TRANSPARENT
		'\x36':	'\u0096',	 // NUMERIC BACKSPACE
		'\x37':	'\u0004',	 // END OF TRANSMISSION
		'\x38':	'\u0098',	 // SUBSCRIPT
		'\x39':	'\u0099',	 // INDENT TABULATION
		'\x3A':	'\u009A',	 // REVERSE FORM FEED
		'\x3B':	'\u009B',	 // CUSTOMER USE THREE
		'\x3C':	'\u0014',	 // DEVICE CONTROL FOUR
		'\x3D':	'\u0015',	 // NEGATIVE ACKNOWLEDGE
		'\x3E':	'\u009E',	 // <reserved>
		'\x3F':	'\u001A',	 // SUBSTITUTE
		'\x40':	'\u0020',	 // SPACE
		'\x41':	'\u05D0',	 // HEBREW LETTER ALEF
		'\x42':	'\u05D1',	 // HEBREW LETTER BET
		'\x43':	'\u05D2',	 // HEBREW LETTER GIMEL
		'\x44':	'\u05D3',	 // HEBREW LETTER DALET
		'\x45':	'\u05D4',	 // HEBREW LETTER HE
		'\x46':	'\u05D5',	 // HEBREW LETTER VAV
		'\x47':	'\u05D6',	 // HEBREW LETTER ZAYIN
		'\x48':	'\u05D7',	 // HEBREW LETTER HET
		'\x49':	'\u05D8',	 // HEBREW LETTER TET
		'\x4A':	'\u00A2',	 // CENT SIGN
		'\x4B':	'\u002E',	 // FULL STOP
		'\x4C':	'\u003C',	 // LESS-THAN SIGN
		'\x4D':	'\u0028',	 // LEFT PARENTHESIS
		'\x4E':	'\u002B',	 // PLUS SIGN
		'\x4F':	'\u007C',	 // VERTICAL LINE
		'\x50':	'\u0026',	 // AMPERSAND
		'\x51':	'\u05D9',	 // HEBREW LETTER YOD
		'\x52':	'\u05DA',	 // HEBREW LETTER FINAL KAF
		'\x53':	'\u05DB',	 // HEBREW LETTER KAF
		'\x54':	'\u05DC',	 // HEBREW LETTER LAMED
		'\x55':	'\u05DD',	 // HEBREW LETTER FINAL MEM
		'\x56':	'\u05DE',	 // HEBREW LETTER MEM
		'\x57':	'\u05DF',	 // HEBREW LETTER FINAL NUN
		'\x58':	'\u05E0',	 // HEBREW LETTER NUN
		'\x59':	'\u05E1',	 // HEBREW LETTER SAMEKH
		'\x5A':	'\u0021',	 // EXCLAMATION MARK
		'\x5B':	'\u0024',	 // DOLLAR SIGN
		'\x5C':	'\u002A',	 // ASTERISK
		'\x5D':	'\u0029',	 // RIGHT PARENTHESIS
		'\x5E':	'\u003B',	 // SEMICOLON
		'\x5F':	'\u00AC',	 // NOT SIGN
		'\x60':	'\u002D',	 // HYPHEN-MINUS
		'\x61':	'\u002F',	 // SOLIDUS
		'\x62':	'\u05E2',	 // HEBREW LETTER AYIN
		'\x63':	'\u05E3',	 // HEBREW LETTER FINAL PE
		'\x64':	'\u05E4',	 // HEBREW LETTER PE
		'\x65':	'\u05E5',	 // HEBREW LETTER FINAL TSADI
		'\x66':	'\u05E6',	 // HEBREW LETTER TSADI
		'\x67':	'\u05E7',	 // HEBREW LETTER QOF
		'\x68':	'\u05E8',	 // HEBREW LETTER RESH
		'\x69':	'\u05E9',	 // HEBREW LETTER SHIN
		'\x6A':	'\u00A6',	 // BROKEN BAR
		'\x6B':	'\u002C',	 // COMMA
		'\x6C':	'\u0025',	 // PERCENT SIGN
		'\x6D':	'\u005F',	 // LOW LINE
		'\x6E':	'\u003E',	 // GREATER-THAN SIGN
		'\x6F':	'\u003F',	 // QUESTION MARK
		// '\x70' UNDEFINED
		'\x71':	'\u05EA',	 // HEBREW LETTER TAV
		// '\x72' UNDEFINED
		// '\x73' UNDEFINED
		'\x74':	'\u00A0',	 // NO-BREAK SPACE
		// '\x75' UNDEFINED
		// '\x76' UNDEFINED
		// '\x77' UNDEFINED
		'\x78':	'\u2017',	 // DOUBLE LOW LINE
		'\x79':	'\u0060',	 // GRAVE ACCENT
		'\x7A':	'\u003A',	 // COLON
		'\x7B':	'\u0023',	 // NUMBER SIGN
		'\x7C':	'\u0040',	 // COMMERCIAL AT
		'\x7D':	'\u0027',	 // APOSTROPHE
		'\x7E':	'\u003D',	 // EQUALS SIGN
		'\x7F':	'\u0022',	 // QUOTATION MARK
		// '\x80' UNDEFINED
		'\x81':	'\u0061',	 // LATIN SMALL LETTER A
		'\x82':	'\u0062',	 // LATIN SMALL LETTER B
		'\x83':	'\u0063',	 // LATIN SMALL LETTER C
		'\x84':	'\u0064',	 // LATIN SMALL LETTER D
		'\x85':	'\u0065',	 // LATIN SMALL LETTER E
		'\x86':	'\u0066',	 // LATIN SMALL LETTER F
		'\x87':	'\u0067',	 // LATIN SMALL LETTER G
		'\x88':	'\u0068',	 // LATIN SMALL LETTER H
		'\x89':	'\u0069',	 // LATIN SMALL LETTER I
		'\x8A':	'\u00AB',	 // LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\x8B':	'\u00BB',	 // RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		// '\x8C' UNDEFINED
		// '\x8D' UNDEFINED
		// '\x8E' UNDEFINED
		'\x8F':	'\u00B1',	 // PLUS-MINUS SIGN
		'\x90':	'\u00B0',	 // DEGREE SIGN
		'\x91':	'\u006A',	 // LATIN SMALL LETTER J
		'\x92':	'\u006B',	 // LATIN SMALL LETTER K
		'\x93':	'\u006C',	 // LATIN SMALL LETTER L
		'\x94':	'\u006D',	 // LATIN SMALL LETTER M
		'\x95':	'\u006E',	 // LATIN SMALL LETTER N
		'\x96':	'\u006F',	 // LATIN SMALL LETTER O
		'\x97':	'\u0070',	 // LATIN SMALL LETTER P
		'\x98':	'\u0071',	 // LATIN SMALL LETTER Q
		'\x99':	'\u0072',	 // LATIN SMALL LETTER R
		// '\x9A' UNDEFINED
		// '\x9B' UNDEFINED
		// '\x9C' UNDEFINED
		'\x9D':	'\u00B8',	 // CEDILLA
		// '\x9E' UNDEFINED
		'\x9F':	'\u00A4',	 // CURRENCY SIGN
		'\xA0':	'\u00B5',	 // MICRO SIGN
		'\xA1':	'\u007E',	 // TILDE
		'\xA2':	'\u0073',	 // LATIN SMALL LETTER S
		'\xA3':	'\u0074',	 // LATIN SMALL LETTER T
		'\xA4':	'\u0075',	 // LATIN SMALL LETTER U
		'\xA5':	'\u0076',	 // LATIN SMALL LETTER V
		'\xA6':	'\u0077',	 // LATIN SMALL LETTER W
		'\xA7':	'\u0078',	 // LATIN SMALL LETTER X
		'\xA8':	'\u0079',	 // LATIN SMALL LETTER Y
		'\xA9':	'\u007A',	 // LATIN SMALL LETTER Z
		// '\xAA' UNDEFINED
		// '\xAB' UNDEFINED
		// '\xAC' UNDEFINED
		// '\xAD' UNDEFINED
		// '\xAE' UNDEFINED
		'\xAF':	'\u00AE',	 // REGISTERED SIGN
		'\xB0':	'\u005E',	 // CIRCUMFLEX ACCENT
		'\xB1':	'\u00A3',	 // POUND SIGN
		'\xB2':	'\u00A5',	 // YEN SIGN
		'\xB3':	'\u00B7',	 // MIDDLE DOT
		'\xB4':	'\u00A9',	 // COPYRIGHT SIGN
		'\xB5':	'\u00A7',	 // SECTION SIGN
		'\xB6':	'\u00B6',	 // PILCROW SIGN
		'\xB7':	'\u00BC',	 // VULGAR FRACTION ONE QUARTER
		'\xB8':	'\u00BD',	 // VULGAR FRACTION ONE HALF
		'\xB9':	'\u00BE',	 // VULGAR FRACTION THREE QUARTERS
		'\xBA':	'\u005B',	 // LEFT SQUARE BRACKET
		'\xBB':	'\u005D',	 // RIGHT SQUARE BRACKET
		'\xBC':	'\u00AF',	 // MACRON
		'\xBD':	'\u00A8',	 // DIAERESIS
		'\xBE':	'\u00B4',	 // ACUTE ACCENT
		'\xBF':	'\u00D7',	 // MULTIPLICATION SIGN
		'\xC0':	'\u007B',	 // LEFT CURLY BRACKET
		'\xC1':	'\u0041',	 // LATIN CAPITAL LETTER A
		'\xC2':	'\u0042',	 // LATIN CAPITAL LETTER B
		'\xC3':	'\u0043',	 // LATIN CAPITAL LETTER C
		'\xC4':	'\u0044',	 // LATIN CAPITAL LETTER D
		'\xC5':	'\u0045',	 // LATIN CAPITAL LETTER E
		'\xC6':	'\u0046',	 // LATIN CAPITAL LETTER F
		'\xC7':	'\u0047',	 // LATIN CAPITAL LETTER G
		'\xC8':	'\u0048',	 // LATIN CAPITAL LETTER H
		'\xC9':	'\u0049',	 // LATIN CAPITAL LETTER I
		'\xCA':	'\u00AD',	 // SOFT HYPHEN
		// '\xCB' UNDEFINED
		// '\xCC' UNDEFINED
		// '\xCD' UNDEFINED
		// '\xCE' UNDEFINED
		// '\xCF' UNDEFINED
		'\xD0':	'\u007D',	 // RIGHT CURLY BRACKET
		'\xD1':	'\u004A',	 // LATIN CAPITAL LETTER J
		'\xD2':	'\u004B',	 // LATIN CAPITAL LETTER K
		'\xD3':	'\u004C',	 // LATIN CAPITAL LETTER L
		'\xD4':	'\u004D',	 // LATIN CAPITAL LETTER M
		'\xD5':	'\u004E',	 // LATIN CAPITAL LETTER N
		'\xD6':	'\u004F',	 // LATIN CAPITAL LETTER O
		'\xD7':	'\u0050',	 // LATIN CAPITAL LETTER P
		'\xD8':	'\u0051',	 // LATIN CAPITAL LETTER Q
		'\xD9':	'\u0052',	 // LATIN CAPITAL LETTER R
		'\xDA':	'\u00B9',	 // SUPERSCRIPT ONE
		// '\xDB' UNDEFINED
		// '\xDC' UNDEFINED
		// '\xDD' UNDEFINED
		// '\xDE' UNDEFINED
		// '\xDF' UNDEFINED
		'\xE0':	'\u005C',	 // REVERSE SOLIDUS
		'\xE1':	'\u00F7',	 // DIVISION SIGN
		'\xE2':	'\u0053',	 // LATIN CAPITAL LETTER S
		'\xE3':	'\u0054',	 // LATIN CAPITAL LETTER T
		'\xE4':	'\u0055',	 // LATIN CAPITAL LETTER U
		'\xE5':	'\u0056',	 // LATIN CAPITAL LETTER V
		'\xE6':	'\u0057',	 // LATIN CAPITAL LETTER W
		'\xE7':	'\u0058',	 // LATIN CAPITAL LETTER X
		'\xE8':	'\u0059',	 // LATIN CAPITAL LETTER Y
		'\xE9':	'\u005A',	 // LATIN CAPITAL LETTER Z
		'\xEA':	'\u00B2',	 // SUPERSCRIPT TWO
		// '\xEB' UNDEFINED
		// '\xEC' UNDEFINED
		// '\xED' UNDEFINED
		// '\xEE' UNDEFINED
		// '\xEF' UNDEFINED
		'\xF0':	'\u0030',	 // DIGIT ZERO
		'\xF1':	'\u0031',	 // DIGIT ONE
		'\xF2':	'\u0032',	 // DIGIT TWO
		'\xF3':	'\u0033',	 // DIGIT THREE
		'\xF4':	'\u0034',	 // DIGIT FOUR
		'\xF5':	'\u0035',	 // DIGIT FIVE
		'\xF6':	'\u0036',	 // DIGIT SIX
		'\xF7':	'\u0037',	 // DIGIT SEVEN
		'\xF8':	'\u0038',	 // DIGIT EIGHT
		'\xF9':	'\u0039',	 // DIGIT NINE
		'\xFA':	'\u00B3',	 // SUPERSCRIPT THREE
		// '\xFB' UNDEFINED
		// '\xFC' UNDEFINED
		// '\xFD' UNDEFINED
		// '\xFE' UNDEFINED
		'\xFF':	'\u009F',	 // EIGHT ONES

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	newCodec := &codecMap8Bit{EncodeMap: charmapEncode, DecodeMap: charmapDecode}

	register(newCodec, "CP424", "CP-424", "424")

}
