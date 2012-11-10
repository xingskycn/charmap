
package charmap

type codecCP1006 struct {
	EncodeMap map[rune]byte
	DecodeMap map[byte]rune
}

func (c *codecCP1006) Encode(s string) (string, error) {
	return mapRunesToBytes(c.EncodeMap, s)
}

func (c *codecCP1006) Decode(s string) (string, error) {
	return mapBytesToRunes(c.DecodeMap, s)
}

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
		'\xA1':	'\u06F0',	 // 	EXTENDED ARABIC-INDIC DIGIT ZERO
		'\xA2':	'\u06F1',	 // 	EXTENDED ARABIC-INDIC DIGIT ONE
		'\xA3':	'\u06F2',	 // 	EXTENDED ARABIC-INDIC DIGIT TWO
		'\xA4':	'\u06F3',	 // 	EXTENDED ARABIC-INDIC DIGIT THREE
		'\xA5':	'\u06F4',	 // 	EXTENDED ARABIC-INDIC DIGIT FOUR
		'\xA6':	'\u06F5',	 // 	EXTENDED ARABIC-INDIC DIGIT FIVE
		'\xA7':	'\u06F6',	 // 	EXTENDED ARABIC-INDIC DIGIT SIX
		'\xA8':	'\u06F7',	 // 	EXTENDED ARABIC-INDIC DIGIT SEVEN
		'\xA9':	'\u06F8',	 // 	EXTENDED ARABIC-INDIC DIGIT EIGHT
		'\xAA':	'\u06F9',	 // 	EXTENDED ARABIC-INDIC DIGIT NINE
		'\xAB':	'\u060C',	 // 	ARABIC COMMA
		'\xAC':	'\u061B',	 // 	ARABIC SEMICOLON
		'\xAD':	'\u00AD',	 // 	SOFT HYPHEN
		'\xAE':	'\u061F',	 // 	ARABIC QUESTION MARK
		'\xAF':	'\uFE81',	 // 	ARABIC LETTER ALEF WITH MADDA ABOVE ISOLATED FORM
		'\xB0':	'\uFE8D',	 // 	ARABIC LETTER ALEF ISOLATED FORM
		'\xB1':	'\uFE8E',	 // 	ARABIC LETTER ALEF FINAL FORM
		'\xB2':	'\uFE8E',	 // 	ARABIC LETTER ALEF FINAL FORM
		'\xB3':	'\uFE8F',	 // 	ARABIC LETTER BEH ISOLATED FORM
		'\xB4':	'\uFE91',	 // 	ARABIC LETTER BEH INITIAL FORM
		'\xB5':	'\uFB56',	 // 	ARABIC LETTER PEH ISOLATED FORM
		'\xB6':	'\uFB58',	 // 	ARABIC LETTER PEH INITIAL FORM
		'\xB7':	'\uFE93',	 // 	ARABIC LETTER TEH MARBUTA ISOLATED FORM
		'\xB8':	'\uFE95',	 // 	ARABIC LETTER TEH ISOLATED FORM
		'\xB9':	'\uFE97',	 // 	ARABIC LETTER TEH INITIAL FORM
		'\xBA':	'\uFB66',	 // 	ARABIC LETTER TTEH ISOLATED FORM
		'\xBB':	'\uFB68',	 // 	ARABIC LETTER TTEH INITIAL FORM
		'\xBC':	'\uFE99',	 // 	ARABIC LETTER THEH ISOLATED FORM
		'\xBD':	'\uFE9B',	 // 	ARABIC LETTER THEH INITIAL FORM
		'\xBE':	'\uFE9D',	 // 	ARABIC LETTER JEEM ISOLATED FORM
		'\xBF':	'\uFE9F',	 // 	ARABIC LETTER JEEM INITIAL FORM
		'\xC0':	'\uFB7A',	 // 	ARABIC LETTER TCHEH ISOLATED FORM
		'\xC1':	'\uFB7C',	 // 	ARABIC LETTER TCHEH INITIAL FORM
		'\xC2':	'\uFEA1',	 // 	ARABIC LETTER HAH ISOLATED FORM
		'\xC3':	'\uFEA3',	 // 	ARABIC LETTER HAH INITIAL FORM
		'\xC4':	'\uFEA5',	 // 	ARABIC LETTER KHAH ISOLATED FORM
		'\xC5':	'\uFEA7',	 // 	ARABIC LETTER KHAH INITIAL FORM
		'\xC6':	'\uFEA9',	 // 	ARABIC LETTER DAL ISOLATED FORM
		'\xC7':	'\uFB84',	 // 	ARABIC LETTER DAHAL ISOLATED FORMN
		'\xC8':	'\uFEAB',	 // 	ARABIC LETTER THAL ISOLATED FORM
		'\xC9':	'\uFEAD',	 // 	ARABIC LETTER REH ISOLATED FORM
		'\xCA':	'\uFB8C',	 // 	ARABIC LETTER RREH ISOLATED FORM
		'\xCB':	'\uFEAF',	 // 	ARABIC LETTER ZAIN ISOLATED FORM
		'\xCC':	'\uFB8A',	 // 	ARABIC LETTER JEH ISOLATED FORM
		'\xCD':	'\uFEB1',	 // 	ARABIC LETTER SEEN ISOLATED FORM
		'\xCE':	'\uFEB3',	 // 	ARABIC LETTER SEEN INITIAL FORM
		'\xCF':	'\uFEB5',	 // 	ARABIC LETTER SHEEN ISOLATED FORM
		'\xD0':	'\uFEB7',	 // 	ARABIC LETTER SHEEN INITIAL FORM
		'\xD1':	'\uFEB9',	 // 	ARABIC LETTER SAD ISOLATED FORM
		'\xD2':	'\uFEBB',	 // 	ARABIC LETTER SAD INITIAL FORM
		'\xD3':	'\uFEBD',	 // 	ARABIC LETTER DAD ISOLATED FORM
		'\xD4':	'\uFEBF',	 // 	ARABIC LETTER DAD INITIAL FORM
		'\xD5':	'\uFEC1',	 // 	ARABIC LETTER TAH ISOLATED FORM
		'\xD6':	'\uFEC5',	 // 	ARABIC LETTER ZAH ISOLATED FORM
		'\xD7':	'\uFEC9',	 // 	ARABIC LETTER AIN ISOLATED FORM
		'\xD8':	'\uFECA',	 // 	ARABIC LETTER AIN FINAL FORM
		'\xD9':	'\uFECB',	 // 	ARABIC LETTER AIN INITIAL FORM
		'\xDA':	'\uFECC',	 // 	ARABIC LETTER AIN MEDIAL FORM
		'\xDB':	'\uFECD',	 // 	ARABIC LETTER GHAIN ISOLATED FORM
		'\xDC':	'\uFECE',	 // 	ARABIC LETTER GHAIN FINAL FORM
		'\xDD':	'\uFECF',	 // 	ARABIC LETTER GHAIN INITIAL FORM
		'\xDE':	'\uFED0',	 // 	ARABIC LETTER GHAIN MEDIAL FORM
		'\xDF':	'\uFED1',	 // 	ARABIC LETTER FEH ISOLATED FORM
		'\xE0':	'\uFED3',	 // 	ARABIC LETTER FEH INITIAL FORM
		'\xE1':	'\uFED5',	 // 	ARABIC LETTER QAF ISOLATED FORM
		'\xE2':	'\uFED7',	 // 	ARABIC LETTER QAF INITIAL FORM
		'\xE3':	'\uFED9',	 // 	ARABIC LETTER KAF ISOLATED FORM
		'\xE4':	'\uFEDB',	 // 	ARABIC LETTER KAF INITIAL FORM
		'\xE5':	'\uFB92',	 // 	ARABIC LETTER GAF ISOLATED FORM
		'\xE6':	'\uFB94',	 // 	ARABIC LETTER GAF INITIAL FORM
		'\xE7':	'\uFEDD',	 // 	ARABIC LETTER LAM ISOLATED FORM
		'\xE8':	'\uFEDF',	 // 	ARABIC LETTER LAM INITIAL FORM
		'\xE9':	'\uFEE0',	 // 	ARABIC LETTER LAM MEDIAL FORM
		'\xEA':	'\uFEE1',	 // 	ARABIC LETTER MEEM ISOLATED FORM
		'\xEB':	'\uFEE3',	 // 	ARABIC LETTER MEEM INITIAL FORM
		'\xEC':	'\uFB9E',	 // 	ARABIC LETTER NOON GHUNNA ISOLATED FORM
		'\xED':	'\uFEE5',	 // 	ARABIC LETTER NOON ISOLATED FORM
		'\xEE':	'\uFEE7',	 // 	ARABIC LETTER NOON INITIAL FORM
		'\xEF':	'\uFE85',	 // 	ARABIC LETTER WAW WITH HAMZA ABOVE ISOLATED FORM
		'\xF0':	'\uFEED',	 // 	ARABIC LETTER WAW ISOLATED FORM
		'\xF1':	'\uFBA6',	 // 	ARABIC LETTER HEH GOAL ISOLATED FORM
		'\xF2':	'\uFBA8',	 // 	ARABIC LETTER HEH GOAL INITIAL FORM
		'\xF3':	'\uFBA9',	 // 	ARABIC LETTER HEH GOAL MEDIAL FORM
		'\xF4':	'\uFBAA',	 // 	ARABIC LETTER HEH DOACHASHMEE ISOLATED FORM
		'\xF5':	'\uFE80',	 // 	ARABIC LETTER HAMZA ISOLATED FORM
		'\xF6':	'\uFE89',	 // 	ARABIC LETTER YEH WITH HAMZA ABOVE ISOLATED FORM
		'\xF7':	'\uFE8A',	 // 	ARABIC LETTER YEH WITH HAMZA ABOVE FINAL FORM
		'\xF8':	'\uFE8B',	 // 	ARABIC LETTER YEH WITH HAMZA ABOVE INITIAL FORM
		'\xF9':	'\uFEF1',	 // 	ARABIC LETTER YEH ISOLATED FORM
		'\xFA':	'\uFEF2',	 // 	ARABIC LETTER YEH FINAL FORM
		'\xFB':	'\uFEF3',	 // 	ARABIC LETTER YEH INITIAL FORM
		'\xFC':	'\uFBB0',	 // 	ARABIC LETTER YEH BARREE WITH HAMZA ABOVE ISOLATED FORM
		'\xFD':	'\uFBAE',	 // 	ARABIC LETTER YEH BARREE ISOLATED FORM
		'\xFE':	'\uFE7C',	 // 	ARABIC SHADDA ISOLATED FORM
		'\xFF':	'\uFE7D',	 // 	ARABIC SHADDA MEDIAL FORM

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	codec := &codecCP1006{
		EncodeMap: charmapEncode,
		DecodeMap: charmapDecode,
	}

	cm := charmap{
		Name: "cp1006",
		Aliases: []string{
			"1006",
		},
		Codec: codec,
	}

	register(cm)
}
