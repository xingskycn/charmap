
package charmap

func init() {

	charmapDecode := map[byte]rune{
		'\x00':	'\u0000',	 // 	NULL
		'\x01':	'\u0001',	 // 	START OF HEADING
		'\x02':	'\u0002',	 // 	START OF TEXT
		'\x03':	'\u0003',	 // 	END OF TEXT
		'\x04':	'\u0004',	 // 	END OF TRANSMISSION
		'\x05':	'\u0005',	 // 	ENQUIRY
		'\x06':	'\u0006',	 // 	ACKNOWLEDGE
		'\x07':	'\u0007',	 // 	BELL
		'\x08':	'\u0008',	 // 	BACKSPACE
		'\x09':	'\u0009',	 // 	HORIZONTAL TABULATION
		'\x0A':	'\u000A',	 // 	LINE FEED
		'\x0B':	'\u000B',	 // 	VERTICAL TABULATION
		'\x0C':	'\u000C',	 // 	FORM FEED
		'\x0D':	'\u000D',	 // 	CARRIAGE RETURN
		'\x0E':	'\u000E',	 // 	SHIFT OUT
		'\x0F':	'\u000F',	 // 	SHIFT IN
		'\x10':	'\u0010',	 // 	DATA LINK ESCAPE
		'\x11':	'\u0011',	 // 	DEVICE CONTROL ONE
		'\x12':	'\u0012',	 // 	DEVICE CONTROL TWO
		'\x13':	'\u0013',	 // 	DEVICE CONTROL THREE
		'\x14':	'\u0014',	 // 	DEVICE CONTROL FOUR
		'\x15':	'\u0015',	 // 	NEGATIVE ACKNOWLEDGE
		'\x16':	'\u0016',	 // 	SYNCHRONOUS IDLE
		'\x17':	'\u0017',	 // 	END OF TRANSMISSION BLOCK
		'\x18':	'\u0018',	 // 	CANCEL
		'\x19':	'\u0019',	 // 	END OF MEDIUM
		'\x1A':	'\u001A',	 // 	SUBSTITUTE
		'\x1B':	'\u001B',	 // 	ESCAPE
		'\x1C':	'\u001C',	 // 	FILE SEPARATOR
		'\x1D':	'\u001D',	 // 	GROUP SEPARATOR
		'\x1E':	'\u001E',	 // 	RECORD SEPARATOR
		'\x1F':	'\u001F',	 // 	UNIT SEPARATOR
		'\x20':	'\u0020',	 // 	SPACE
		'\x21':	'\u0021',	 // 	EXCLAMATION MARK
		'\x22':	'\u0022',	 // 	QUOTATION MARK
		'\x23':	'\u0023',	 // 	NUMBER SIGN
		'\x24':	'\u0024',	 // 	DOLLAR SIGN
		'\x25':	'\u0025',	 // 	PERCENT SIGN
		'\x26':	'\u0026',	 // 	AMPERSAND
		'\x27':	'\u0027',	 // 	APOSTROPHE
		'\x28':	'\u0028',	 // 	LEFT PARENTHESIS
		'\x29':	'\u0029',	 // 	RIGHT PARENTHESIS
		'\x2A':	'\u002A',	 // 	ASTERISK
		'\x2B':	'\u002B',	 // 	PLUS SIGN
		'\x2C':	'\u002C',	 // 	COMMA
		'\x2D':	'\u002D',	 // 	HYPHEN-MINUS
		'\x2E':	'\u002E',	 // 	FULL STOP
		'\x2F':	'\u002F',	 // 	SOLIDUS
		'\x30':	'\u0030',	 // 	DIGIT ZERO
		'\x31':	'\u0031',	 // 	DIGIT ONE
		'\x32':	'\u0032',	 // 	DIGIT TWO
		'\x33':	'\u0033',	 // 	DIGIT THREE
		'\x34':	'\u0034',	 // 	DIGIT FOUR
		'\x35':	'\u0035',	 // 	DIGIT FIVE
		'\x36':	'\u0036',	 // 	DIGIT SIX
		'\x37':	'\u0037',	 // 	DIGIT SEVEN
		'\x38':	'\u0038',	 // 	DIGIT EIGHT
		'\x39':	'\u0039',	 // 	DIGIT NINE
		'\x3A':	'\u003A',	 // 	COLON
		'\x3B':	'\u003B',	 // 	SEMICOLON
		'\x3C':	'\u003C',	 // 	LESS-THAN SIGN
		'\x3D':	'\u003D',	 // 	EQUALS SIGN
		'\x3E':	'\u003E',	 // 	GREATER-THAN SIGN
		'\x3F':	'\u003F',	 // 	QUESTION MARK
		'\x40':	'\u0040',	 // 	COMMERCIAL AT
		'\x41':	'\u0041',	 // 	LATIN CAPITAL LETTER A
		'\x42':	'\u0042',	 // 	LATIN CAPITAL LETTER B
		'\x43':	'\u0043',	 // 	LATIN CAPITAL LETTER C
		'\x44':	'\u0044',	 // 	LATIN CAPITAL LETTER D
		'\x45':	'\u0045',	 // 	LATIN CAPITAL LETTER E
		'\x46':	'\u0046',	 // 	LATIN CAPITAL LETTER F
		'\x47':	'\u0047',	 // 	LATIN CAPITAL LETTER G
		'\x48':	'\u0048',	 // 	LATIN CAPITAL LETTER H
		'\x49':	'\u0049',	 // 	LATIN CAPITAL LETTER I
		'\x4A':	'\u004A',	 // 	LATIN CAPITAL LETTER J
		'\x4B':	'\u004B',	 // 	LATIN CAPITAL LETTER K
		'\x4C':	'\u004C',	 // 	LATIN CAPITAL LETTER L
		'\x4D':	'\u004D',	 // 	LATIN CAPITAL LETTER M
		'\x4E':	'\u004E',	 // 	LATIN CAPITAL LETTER N
		'\x4F':	'\u004F',	 // 	LATIN CAPITAL LETTER O
		'\x50':	'\u0050',	 // 	LATIN CAPITAL LETTER P
		'\x51':	'\u0051',	 // 	LATIN CAPITAL LETTER Q
		'\x52':	'\u0052',	 // 	LATIN CAPITAL LETTER R
		'\x53':	'\u0053',	 // 	LATIN CAPITAL LETTER S
		'\x54':	'\u0054',	 // 	LATIN CAPITAL LETTER T
		'\x55':	'\u0055',	 // 	LATIN CAPITAL LETTER U
		'\x56':	'\u0056',	 // 	LATIN CAPITAL LETTER V
		'\x57':	'\u0057',	 // 	LATIN CAPITAL LETTER W
		'\x58':	'\u0058',	 // 	LATIN CAPITAL LETTER X
		'\x59':	'\u0059',	 // 	LATIN CAPITAL LETTER Y
		'\x5A':	'\u005A',	 // 	LATIN CAPITAL LETTER Z
		'\x5B':	'\u005B',	 // 	LEFT SQUARE BRACKET
		'\x5C':	'\u005C',	 // 	REVERSE SOLIDUS
		'\x5D':	'\u005D',	 // 	RIGHT SQUARE BRACKET
		'\x5E':	'\u005E',	 // 	CIRCUMFLEX ACCENT
		'\x5F':	'\u005F',	 // 	LOW LINE
		'\x60':	'\u0060',	 // 	GRAVE ACCENT
		'\x61':	'\u0061',	 // 	LATIN SMALL LETTER A
		'\x62':	'\u0062',	 // 	LATIN SMALL LETTER B
		'\x63':	'\u0063',	 // 	LATIN SMALL LETTER C
		'\x64':	'\u0064',	 // 	LATIN SMALL LETTER D
		'\x65':	'\u0065',	 // 	LATIN SMALL LETTER E
		'\x66':	'\u0066',	 // 	LATIN SMALL LETTER F
		'\x67':	'\u0067',	 // 	LATIN SMALL LETTER G
		'\x68':	'\u0068',	 // 	LATIN SMALL LETTER H
		'\x69':	'\u0069',	 // 	LATIN SMALL LETTER I
		'\x6A':	'\u006A',	 // 	LATIN SMALL LETTER J
		'\x6B':	'\u006B',	 // 	LATIN SMALL LETTER K
		'\x6C':	'\u006C',	 // 	LATIN SMALL LETTER L
		'\x6D':	'\u006D',	 // 	LATIN SMALL LETTER M
		'\x6E':	'\u006E',	 // 	LATIN SMALL LETTER N
		'\x6F':	'\u006F',	 // 	LATIN SMALL LETTER O
		'\x70':	'\u0070',	 // 	LATIN SMALL LETTER P
		'\x71':	'\u0071',	 // 	LATIN SMALL LETTER Q
		'\x72':	'\u0072',	 // 	LATIN SMALL LETTER R
		'\x73':	'\u0073',	 // 	LATIN SMALL LETTER S
		'\x74':	'\u0074',	 // 	LATIN SMALL LETTER T
		'\x75':	'\u0075',	 // 	LATIN SMALL LETTER U
		'\x76':	'\u0076',	 // 	LATIN SMALL LETTER V
		'\x77':	'\u0077',	 // 	LATIN SMALL LETTER W
		'\x78':	'\u0078',	 // 	LATIN SMALL LETTER X
		'\x79':	'\u0079',	 // 	LATIN SMALL LETTER Y
		'\x7A':	'\u007A',	 // 	LATIN SMALL LETTER Z
		'\x7B':	'\u007B',	 // 	LEFT CURLY BRACKET
		'\x7C':	'\u007C',	 // 	VERTICAL LINE
		'\x7D':	'\u007D',	 // 	RIGHT CURLY BRACKET
		'\x7E':	'\u007E',	 // 	TILDE
		'\x7F':	'\u007F',	 // 	DELETE
		'\x80':	'\u0080',	 // 	<control>
		'\x81':	'\u0081',	 // 	<control>
		'\x82':	'\u0082',	 // 	<control>
		'\x83':	'\u0083',	 // 	<control>
		'\x84':	'\u0084',	 // 	<control>
		'\x85':	'\u0085',	 // 	<control>
		'\x86':	'\u0086',	 // 	<control>
		'\x87':	'\u0087',	 // 	<control>
		'\x88':	'\u0088',	 // 	<control>
		'\x89':	'\u0089',	 // 	<control>
		'\x8A':	'\u008A',	 // 	<control>
		'\x8B':	'\u008B',	 // 	<control>
		'\x8C':	'\u008C',	 // 	<control>
		'\x8D':	'\u008D',	 // 	<control>
		'\x8E':	'\u008E',	 // 	<control>
		'\x8F':	'\u008F',	 // 	<control>
		'\x90':	'\u0090',	 // 	<control>
		'\x91':	'\u0091',	 // 	<control>
		'\x92':	'\u0092',	 // 	<control>
		'\x93':	'\u0093',	 // 	<control>
		'\x94':	'\u0094',	 // 	<control>
		'\x95':	'\u0095',	 // 	<control>
		'\x96':	'\u0096',	 // 	<control>
		'\x97':	'\u0097',	 // 	<control>
		'\x98':	'\u0098',	 // 	<control>
		'\x99':	'\u0099',	 // 	<control>
		'\x9A':	'\u009A',	 // 	<control>
		'\x9B':	'\u009B',	 // 	<control>
		'\x9C':	'\u009C',	 // 	<control>
		'\x9D':	'\u009D',	 // 	<control>
		'\x9E':	'\u009E',	 // 	<control>
		'\x9F':	'\u009F',	 // 	<control>
		'\xA0':	'\u00A0',	 // 	NO-BREAK SPACE
		'\xA1':	'\u2018',	 // 	LEFT SINGLE QUOTATION MARK
		'\xA2':	'\u2019',	 // 	RIGHT SINGLE QUOTATION MARK
		'\xA3':	'\u00A3',	 // 	POUND SIGN
		'\xA4':	'\u20AC',	 // 	EURO SIGN
		'\xA5':	'\u20AF',	 // 	DRACHMA SIGN
		'\xA6':	'\u00A6',	 // 	BROKEN BAR
		'\xA7':	'\u00A7',	 // 	SECTION SIGN
		'\xA8':	'\u00A8',	 // 	DIAERESIS
		'\xA9':	'\u00A9',	 // 	COPYRIGHT SIGN
		'\xAA':	'\u037A',	 // 	GREEK YPOGEGRAMMENI
		'\xAB':	'\u00AB',	 // 	LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xAC':	'\u00AC',	 // 	NOT SIGN
		'\xAD':	'\u00AD',	 // 	SOFT HYPHEN
		'\xAF':	'\u2015',	 // 	HORIZONTAL BAR
		'\xB0':	'\u00B0',	 // 	DEGREE SIGN
		'\xB1':	'\u00B1',	 // 	PLUS-MINUS SIGN
		'\xB2':	'\u00B2',	 // 	SUPERSCRIPT TWO
		'\xB3':	'\u00B3',	 // 	SUPERSCRIPT THREE
		'\xB4':	'\u0384',	 // 	GREEK TONOS
		'\xB5':	'\u0385',	 // 	GREEK DIALYTIKA TONOS
		'\xB6':	'\u0386',	 // 	GREEK CAPITAL LETTER ALPHA WITH TONOS
		'\xB7':	'\u00B7',	 // 	MIDDLE DOT
		'\xB8':	'\u0388',	 // 	GREEK CAPITAL LETTER EPSILON WITH TONOS
		'\xB9':	'\u0389',	 // 	GREEK CAPITAL LETTER ETA WITH TONOS
		'\xBA':	'\u038A',	 // 	GREEK CAPITAL LETTER IOTA WITH TONOS
		'\xBB':	'\u00BB',	 // 	RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
		'\xBC':	'\u038C',	 // 	GREEK CAPITAL LETTER OMICRON WITH TONOS
		'\xBD':	'\u00BD',	 // 	VULGAR FRACTION ONE HALF
		'\xBE':	'\u038E',	 // 	GREEK CAPITAL LETTER UPSILON WITH TONOS
		'\xBF':	'\u038F',	 // 	GREEK CAPITAL LETTER OMEGA WITH TONOS
		'\xC0':	'\u0390',	 // 	GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS
		'\xC1':	'\u0391',	 // 	GREEK CAPITAL LETTER ALPHA
		'\xC2':	'\u0392',	 // 	GREEK CAPITAL LETTER BETA
		'\xC3':	'\u0393',	 // 	GREEK CAPITAL LETTER GAMMA
		'\xC4':	'\u0394',	 // 	GREEK CAPITAL LETTER DELTA
		'\xC5':	'\u0395',	 // 	GREEK CAPITAL LETTER EPSILON
		'\xC6':	'\u0396',	 // 	GREEK CAPITAL LETTER ZETA
		'\xC7':	'\u0397',	 // 	GREEK CAPITAL LETTER ETA
		'\xC8':	'\u0398',	 // 	GREEK CAPITAL LETTER THETA
		'\xC9':	'\u0399',	 // 	GREEK CAPITAL LETTER IOTA
		'\xCA':	'\u039A',	 // 	GREEK CAPITAL LETTER KAPPA
		'\xCB':	'\u039B',	 // 	GREEK CAPITAL LETTER LAMDA
		'\xCC':	'\u039C',	 // 	GREEK CAPITAL LETTER MU
		'\xCD':	'\u039D',	 // 	GREEK CAPITAL LETTER NU
		'\xCE':	'\u039E',	 // 	GREEK CAPITAL LETTER XI
		'\xCF':	'\u039F',	 // 	GREEK CAPITAL LETTER OMICRON
		'\xD0':	'\u03A0',	 // 	GREEK CAPITAL LETTER PI
		'\xD1':	'\u03A1',	 // 	GREEK CAPITAL LETTER RHO
		'\xD3':	'\u03A3',	 // 	GREEK CAPITAL LETTER SIGMA
		'\xD4':	'\u03A4',	 // 	GREEK CAPITAL LETTER TAU
		'\xD5':	'\u03A5',	 // 	GREEK CAPITAL LETTER UPSILON
		'\xD6':	'\u03A6',	 // 	GREEK CAPITAL LETTER PHI
		'\xD7':	'\u03A7',	 // 	GREEK CAPITAL LETTER CHI
		'\xD8':	'\u03A8',	 // 	GREEK CAPITAL LETTER PSI
		'\xD9':	'\u03A9',	 // 	GREEK CAPITAL LETTER OMEGA
		'\xDA':	'\u03AA',	 // 	GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
		'\xDB':	'\u03AB',	 // 	GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
		'\xDC':	'\u03AC',	 // 	GREEK SMALL LETTER ALPHA WITH TONOS
		'\xDD':	'\u03AD',	 // 	GREEK SMALL LETTER EPSILON WITH TONOS
		'\xDE':	'\u03AE',	 // 	GREEK SMALL LETTER ETA WITH TONOS
		'\xDF':	'\u03AF',	 // 	GREEK SMALL LETTER IOTA WITH TONOS
		'\xE0':	'\u03B0',	 // 	GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS
		'\xE1':	'\u03B1',	 // 	GREEK SMALL LETTER ALPHA
		'\xE2':	'\u03B2',	 // 	GREEK SMALL LETTER BETA
		'\xE3':	'\u03B3',	 // 	GREEK SMALL LETTER GAMMA
		'\xE4':	'\u03B4',	 // 	GREEK SMALL LETTER DELTA
		'\xE5':	'\u03B5',	 // 	GREEK SMALL LETTER EPSILON
		'\xE6':	'\u03B6',	 // 	GREEK SMALL LETTER ZETA
		'\xE7':	'\u03B7',	 // 	GREEK SMALL LETTER ETA
		'\xE8':	'\u03B8',	 // 	GREEK SMALL LETTER THETA
		'\xE9':	'\u03B9',	 // 	GREEK SMALL LETTER IOTA
		'\xEA':	'\u03BA',	 // 	GREEK SMALL LETTER KAPPA
		'\xEB':	'\u03BB',	 // 	GREEK SMALL LETTER LAMDA
		'\xEC':	'\u03BC',	 // 	GREEK SMALL LETTER MU
		'\xED':	'\u03BD',	 // 	GREEK SMALL LETTER NU
		'\xEE':	'\u03BE',	 // 	GREEK SMALL LETTER XI
		'\xEF':	'\u03BF',	 // 	GREEK SMALL LETTER OMICRON
		'\xF0':	'\u03C0',	 // 	GREEK SMALL LETTER PI
		'\xF1':	'\u03C1',	 // 	GREEK SMALL LETTER RHO
		'\xF2':	'\u03C2',	 // 	GREEK SMALL LETTER FINAL SIGMA
		'\xF3':	'\u03C3',	 // 	GREEK SMALL LETTER SIGMA
		'\xF4':	'\u03C4',	 // 	GREEK SMALL LETTER TAU
		'\xF5':	'\u03C5',	 // 	GREEK SMALL LETTER UPSILON
		'\xF6':	'\u03C6',	 // 	GREEK SMALL LETTER PHI
		'\xF7':	'\u03C7',	 // 	GREEK SMALL LETTER CHI
		'\xF8':	'\u03C8',	 // 	GREEK SMALL LETTER PSI
		'\xF9':	'\u03C9',	 // 	GREEK SMALL LETTER OMEGA
		'\xFA':	'\u03CA',	 // 	GREEK SMALL LETTER IOTA WITH DIALYTIKA
		'\xFB':	'\u03CB',	 // 	GREEK SMALL LETTER UPSILON WITH DIALYTIKA
		'\xFC':	'\u03CC',	 // 	GREEK SMALL LETTER OMICRON WITH TONOS
		'\xFD':	'\u03CD',	 // 	GREEK SMALL LETTER UPSILON WITH TONOS
		'\xFE':	'\u03CE',	 // 	GREEK SMALL LETTER OMEGA WITH TONOS

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	newCodec := &codecMap8Bit{EncodeMap: charmapEncode, DecodeMap: charmapDecode}

	register(newCodec, "ISO-8859-7", "8859-7", "ISO8859-7")

}
