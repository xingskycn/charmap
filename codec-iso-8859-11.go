
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
		'\xA1':	'\u0E01',	 // 	THAI CHARACTER KO KAI
		'\xA2':	'\u0E02',	 // 	THAI CHARACTER KHO KHAI
		'\xA3':	'\u0E03',	 // 	THAI CHARACTER KHO KHUAT
		'\xA4':	'\u0E04',	 // 	THAI CHARACTER KHO KHWAI
		'\xA5':	'\u0E05',	 // 	THAI CHARACTER KHO KHON
		'\xA6':	'\u0E06',	 // 	THAI CHARACTER KHO RAKHANG
		'\xA7':	'\u0E07',	 // 	THAI CHARACTER NGO NGU
		'\xA8':	'\u0E08',	 // 	THAI CHARACTER CHO CHAN
		'\xA9':	'\u0E09',	 // 	THAI CHARACTER CHO CHING
		'\xAA':	'\u0E0A',	 // 	THAI CHARACTER CHO CHANG
		'\xAB':	'\u0E0B',	 // 	THAI CHARACTER SO SO
		'\xAC':	'\u0E0C',	 // 	THAI CHARACTER CHO CHOE
		'\xAD':	'\u0E0D',	 // 	THAI CHARACTER YO YING
		'\xAE':	'\u0E0E',	 // 	THAI CHARACTER DO CHADA
		'\xAF':	'\u0E0F',	 // 	THAI CHARACTER TO PATAK
		'\xB0':	'\u0E10',	 // 	THAI CHARACTER THO THAN
		'\xB1':	'\u0E11',	 // 	THAI CHARACTER THO NANGMONTHO
		'\xB2':	'\u0E12',	 // 	THAI CHARACTER THO PHUTHAO
		'\xB3':	'\u0E13',	 // 	THAI CHARACTER NO NEN
		'\xB4':	'\u0E14',	 // 	THAI CHARACTER DO DEK
		'\xB5':	'\u0E15',	 // 	THAI CHARACTER TO TAO
		'\xB6':	'\u0E16',	 // 	THAI CHARACTER THO THUNG
		'\xB7':	'\u0E17',	 // 	THAI CHARACTER THO THAHAN
		'\xB8':	'\u0E18',	 // 	THAI CHARACTER THO THONG
		'\xB9':	'\u0E19',	 // 	THAI CHARACTER NO NU
		'\xBA':	'\u0E1A',	 // 	THAI CHARACTER BO BAIMAI
		'\xBB':	'\u0E1B',	 // 	THAI CHARACTER PO PLA
		'\xBC':	'\u0E1C',	 // 	THAI CHARACTER PHO PHUNG
		'\xBD':	'\u0E1D',	 // 	THAI CHARACTER FO FA
		'\xBE':	'\u0E1E',	 // 	THAI CHARACTER PHO PHAN
		'\xBF':	'\u0E1F',	 // 	THAI CHARACTER FO FAN
		'\xC0':	'\u0E20',	 // 	THAI CHARACTER PHO SAMPHAO
		'\xC1':	'\u0E21',	 // 	THAI CHARACTER MO MA
		'\xC2':	'\u0E22',	 // 	THAI CHARACTER YO YAK
		'\xC3':	'\u0E23',	 // 	THAI CHARACTER RO RUA
		'\xC4':	'\u0E24',	 // 	THAI CHARACTER RU
		'\xC5':	'\u0E25',	 // 	THAI CHARACTER LO LING
		'\xC6':	'\u0E26',	 // 	THAI CHARACTER LU
		'\xC7':	'\u0E27',	 // 	THAI CHARACTER WO WAEN
		'\xC8':	'\u0E28',	 // 	THAI CHARACTER SO SALA
		'\xC9':	'\u0E29',	 // 	THAI CHARACTER SO RUSI
		'\xCA':	'\u0E2A',	 // 	THAI CHARACTER SO SUA
		'\xCB':	'\u0E2B',	 // 	THAI CHARACTER HO HIP
		'\xCC':	'\u0E2C',	 // 	THAI CHARACTER LO CHULA
		'\xCD':	'\u0E2D',	 // 	THAI CHARACTER O ANG
		'\xCE':	'\u0E2E',	 // 	THAI CHARACTER HO NOKHUK
		'\xCF':	'\u0E2F',	 // 	THAI CHARACTER PAIYANNOI
		'\xD0':	'\u0E30',	 // 	THAI CHARACTER SARA A
		'\xD1':	'\u0E31',	 // 	THAI CHARACTER MAI HAN-AKAT
		'\xD2':	'\u0E32',	 // 	THAI CHARACTER SARA AA
		'\xD3':	'\u0E33',	 // 	THAI CHARACTER SARA AM
		'\xD4':	'\u0E34',	 // 	THAI CHARACTER SARA I
		'\xD5':	'\u0E35',	 // 	THAI CHARACTER SARA II
		'\xD6':	'\u0E36',	 // 	THAI CHARACTER SARA UE
		'\xD7':	'\u0E37',	 // 	THAI CHARACTER SARA UEE
		'\xD8':	'\u0E38',	 // 	THAI CHARACTER SARA U
		'\xD9':	'\u0E39',	 // 	THAI CHARACTER SARA UU
		'\xDA':	'\u0E3A',	 // 	THAI CHARACTER PHINTHU
		'\xDF':	'\u0E3F',	 // 	THAI CURRENCY SYMBOL BAHT
		'\xE0':	'\u0E40',	 // 	THAI CHARACTER SARA E
		'\xE1':	'\u0E41',	 // 	THAI CHARACTER SARA AE
		'\xE2':	'\u0E42',	 // 	THAI CHARACTER SARA O
		'\xE3':	'\u0E43',	 // 	THAI CHARACTER SARA AI MAIMUAN
		'\xE4':	'\u0E44',	 // 	THAI CHARACTER SARA AI MAIMALAI
		'\xE5':	'\u0E45',	 // 	THAI CHARACTER LAKKHANGYAO
		'\xE6':	'\u0E46',	 // 	THAI CHARACTER MAIYAMOK
		'\xE7':	'\u0E47',	 // 	THAI CHARACTER MAITAIKHU
		'\xE8':	'\u0E48',	 // 	THAI CHARACTER MAI EK
		'\xE9':	'\u0E49',	 // 	THAI CHARACTER MAI THO
		'\xEA':	'\u0E4A',	 // 	THAI CHARACTER MAI TRI
		'\xEB':	'\u0E4B',	 // 	THAI CHARACTER MAI CHATTAWA
		'\xEC':	'\u0E4C',	 // 	THAI CHARACTER THANTHAKHAT
		'\xED':	'\u0E4D',	 // 	THAI CHARACTER NIKHAHIT
		'\xEE':	'\u0E4E',	 // 	THAI CHARACTER YAMAKKAN
		'\xEF':	'\u0E4F',	 // 	THAI CHARACTER FONGMAN
		'\xF0':	'\u0E50',	 // 	THAI DIGIT ZERO
		'\xF1':	'\u0E51',	 // 	THAI DIGIT ONE
		'\xF2':	'\u0E52',	 // 	THAI DIGIT TWO
		'\xF3':	'\u0E53',	 // 	THAI DIGIT THREE
		'\xF4':	'\u0E54',	 // 	THAI DIGIT FOUR
		'\xF5':	'\u0E55',	 // 	THAI DIGIT FIVE
		'\xF6':	'\u0E56',	 // 	THAI DIGIT SIX
		'\xF7':	'\u0E57',	 // 	THAI DIGIT SEVEN
		'\xF8':	'\u0E58',	 // 	THAI DIGIT EIGHT
		'\xF9':	'\u0E59',	 // 	THAI DIGIT NINE
		'\xFA':	'\u0E5A',	 // 	THAI CHARACTER ANGKHANKHU
		'\xFB':	'\u0E5B',	 // 	THAI CHARACTER KHOMUT

	}

	charmapEncode := reverseByteRuneMap(charmapDecode)

	newCodec := &codecMap8Bit{EncodeMap: charmapEncode, DecodeMap: charmapDecode}

	register(newCodec, "ISO-8859-11", "8859-11", "ISO8859-11")

}
