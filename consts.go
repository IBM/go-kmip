package kmip

const (
	// KMIP 1.0
	ACTIVATION_DATE                        Tag = 0x420001
	APPLICATION_DATA                       Tag = 0x420002
	APPLICATION_NAMESPACE                  Tag = 0x420003
	APPLICATION_SPECIFIC_INFORMATION       Tag = 0x420004
	ARCHIVE_DATE                           Tag = 0x420005
	ASYNCHRONOUS_CORRELATION_VALUE         Tag = 0x420006
	ASYNCHRONOUS_INDICATOR                 Tag = 0x420007
	ATTRIBUTE                              Tag = 0x420008
	ATTRIBUTE_INDEX                        Tag = 0x420009 // Designated '(Reserved)' in KMIP 2.0
	ATTRIBUTE_NAME                         Tag = 0x42000A
	ATTRIBUTE_VALUE                        Tag = 0x42000B
	AUTHENTICATION                         Tag = 0x42000C
	BATCH_COUNT                            Tag = 0x42000D
	BATCH_ERROR_CONTINUATION_OPTION        Tag = 0x42000E
	BATCH_ITEM                             Tag = 0x42000F
	BATCH_ORDER_OPTION                     Tag = 0x420010
	BLOCK_CIPHER_MODE                      Tag = 0x420011
	CANCELLATION_RESULT                    Tag = 0x420012
	CERTIFICATE                            Tag = 0x420013
	CERTIFICATE_IDENTIFIER                 Tag = 0x420014 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER                     Tag = 0x420015 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER_ALTERNATIVE_NAME    Tag = 0x420016 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_ISSUER_DISTINGUISHED_NAME  Tag = 0x420017 // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_REQUEST                    Tag = 0x420018
	CERTIFICATE_REQUEST_TYPE               Tag = 0x420019
	CERTIFICATE_SUBJECT                    Tag = 0x42001A // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_SUBJECT_ALTERNATIVE_NAME   Tag = 0x42001B // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_SUBJECT_DISTINGUISHED_NAME Tag = 0x42001C // Deprecated, designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_TYPE                       Tag = 0x42001D
	CERTIFICATE_VALUE                      Tag = 0x42001E
	COMMON_TEMPLATE_ATTRIBUTE              Tag = 0x42001F // Designated '(Reserved)' in KMIP 2.0
	COMPROMISE_DATE                        Tag = 0x420020
	COMPROMISE_OCCURRENCE_DATE             Tag = 0x420021
	CONTACT_INFORMATION                    Tag = 0x420022
	CREDENTIAL                             Tag = 0x420023
	CREDENTIAL_TYPE                        Tag = 0x420024
	CREDENTIAL_VALUE                       Tag = 0x420025
	CRITICALITY_INDICATOR                  Tag = 0x420026
	CRT_COEFFICIENT                        Tag = 0x420027
	CRYPTOGRAPHIC_ALGORITHM                Tag = 0x420028
	CRYPTOGRAPHIC_DOMAIN_PARAMETERS        Tag = 0x420029
	CRYPTOGRAPHIC_LENGTH                   Tag = 0x42002A
	CRYPTOGRAPHIC_PARAMETERS               Tag = 0x42002B
	CRYPTOGRAPHIC_USAGE_MASK               Tag = 0x42002C
	CUSTOM_ATTRIBUTE                       Tag = 0x42002D // Designated '(Reserved)' in KMIP 2.0
	D                                      Tag = 0x42002E
	DEACTIVATION_DATE                      Tag = 0x42002F
	DERIVATION_DATA                        Tag = 0x420030
	DERIVATION_METHOD                      Tag = 0x420031
	DERIVATION_PARAMETERS                  Tag = 0x420032
	DESTROY_DATE                           Tag = 0x420033
	DIGEST                                 Tag = 0x420034
	DIGEST_VALUE                           Tag = 0x420035
	ENCRYPTION_KEY_INFORMATION             Tag = 0x420036
	G                                      Tag = 0x420037
	HASHING_ALGORITHM                      Tag = 0x420038
	INITIAL_DATE                           Tag = 0x420039
	INITIALIZATION_VECTOR                  Tag = 0x42003A
	ISSUER                                 Tag = 0x42003B // Deprecated, designated '(Reserved)' in KMIP 2.0
	ITERATION_COUNT                        Tag = 0x42003C
	IV_COUNTER_NONCE                       Tag = 0x42003D
	J                                      Tag = 0x42003E
	KEY                                    Tag = 0x42003F
	KEY_BLOCK                              Tag = 0x420040
	KEY_COMPRESSION_TYPE                   Tag = 0x420041
	KEY_FORMAT_TYPE                        Tag = 0x420042
	KEY_MATERIAL                           Tag = 0x420043
	KEY_PART_IDENTIFIER                    Tag = 0x420044
	KEY_VALUE                              Tag = 0x420045
	KEY_WRAPPING_DATA                      Tag = 0x420046
	KEY_WRAPPING_SPECIFICATION             Tag = 0x420047
	LAST_CHANGE_DATE                       Tag = 0x420048
	LEASE_TIME                             Tag = 0x420049
	LINK                                   Tag = 0x42004A
	LINK_TYPE                              Tag = 0x42004B
	LINKED_OBJECT_IDENTIFIER               Tag = 0x42004C
	MAC_SIGNATURE                          Tag = 0x42004D
	MAC_SIGNATURE_KEY_INFORMATION          Tag = 0x42004E
	MAXIMUM_ITEMS                          Tag = 0x42004F
	MAXIMUM_RESPONSE_SIZE                  Tag = 0x420050
	MESSAGE_EXTENSION                      Tag = 0x420051
	MODULUS                                Tag = 0x420052
	NAME                                   Tag = 0x420053
	NAME_TYPE                              Tag = 0x420054
	NAME_VALUE                             Tag = 0x420055
	OBJECT_GROUP                           Tag = 0x420056
	OBJECT_TYPE                            Tag = 0x420057
	OFFSET                                 Tag = 0x420058
	OPAQUE_DATA_TYPE                       Tag = 0x420059
	OPAQUE_DATA_VALUE                      Tag = 0x42005A
	OPAQUE_OBJECT                          Tag = 0x42005B
	OPERATION                              Tag = 0x42005C
	OPERATION_POLICY_NAME                  Tag = 0x42005D // Designated '(Reserved)' in KMIP 2.0
	P                                      Tag = 0x42005E
	PADDING_METHOD                         Tag = 0x42005F
	PRIME_EXPONENT_P                       Tag = 0x420060
	PRIME_EXPONENT_Q                       Tag = 0x420061
	PRIME_FIELD_SIZE                       Tag = 0x420062
	PRIVATE_EXPONENT                       Tag = 0x420063
	PRIVATE_KEY                            Tag = 0x420064
	PRIVATE_KEY_TEMPLATE_ATTRIBUTE         Tag = 0x420065 // Designated '(Reserved)' in KMIP 2.0
	PRIVATE_KEY_UNIQUE_IDENTIFIER          Tag = 0x420066
	PROCESS_START_DATE                     Tag = 0x420067
	PROTECT_STOP_DATE                      Tag = 0x420068
	PROTOCOL_VERSION                       Tag = 0x420069
	PROTOCOL_VERSION_MAJOR                 Tag = 0x42006A
	PROTOCOL_VERSION_MINOR                 Tag = 0x42006B
	PUBLIC_EXPONENT                        Tag = 0x42006C
	PUBLIC_KEY                             Tag = 0x42006D
	PUBLIC_KEY_TEMPLATE_ATTRIBUTE          Tag = 0x42006E // Designated '(Reserved)' in KMIP 2.0
	PUBLIC_KEY_UNIQUE_IDENTIFIER           Tag = 0x42006F
	PUT_FUNCTION                           Tag = 0x420070
	Q                                      Tag = 0x420071
	Q_STRING                               Tag = 0x420072
	QLENGTH                                Tag = 0x420073
	QUERY_FUNCTION                         Tag = 0x420074
	RECOMMENDED_CURVE                      Tag = 0x420075
	REPLACED_UNIQUE_IDENTIFIER             Tag = 0x420076
	REQUEST_BATCH_ITEM                     Tag = 0x42000F
	REQUEST_HEADER                         Tag = 0x420077
	REQUEST_MESSAGE                        Tag = 0x420078
	REQUEST_PAYLOAD                        Tag = 0x420079
	RESPONSE_BATCH_ITEM                    Tag = 0x42000F
	RESPONSE_HEADER                        Tag = 0x42007A
	RESPONSE_MESSAGE                       Tag = 0x42007B
	RESPONSE_PAYLOAD                       Tag = 0x42007C
	RESULT_MESSAGE                         Tag = 0x42007D
	RESULT_REASON                          Tag = 0x42007E
	RESULT_STATUS                          Tag = 0x42007F
	REVOCATION_MESSAGE                     Tag = 0x420080
	REVOCATION_REASON                      Tag = 0x420081
	REVOCATION_REASON_CODE                 Tag = 0x420082
	KEY_ROLE_TYPE                          Tag = 0x420083
	SALT                                   Tag = 0x420084
	SECRET_DATA                            Tag = 0x420085
	SECRET_DATA_TYPE                       Tag = 0x420086
	SERIAL_NUMBER                          Tag = 0x420087 // Deprecated, designated '(Reserved)' in KMIP 2.0
	SERVER_INFORMATION                     Tag = 0x420088
	SPLIT_KEY                              Tag = 0x420089
	SPLIT_KEY_METHOD                       Tag = 0x42008A
	SPLIT_KEY_PARTS                        Tag = 0x42008B
	SPLIT_KEY_THRESHOLD                    Tag = 0x42008C
	STATE                                  Tag = 0x42008D
	STORAGE_STATUS_MASK                    Tag = 0x42008E
	SYMMETRIC_KEY                          Tag = 0x42008F
	TEMPLATE                               Tag = 0x420090 // Designated '(Reserved)' in KMIP 2.0
	TEMPLATE_ATTRIBUTE                     Tag = 0x420091 // Designated '(Reserved)' in KMIP 2.0
	TIME_STAMP                             Tag = 0x420092
	UNIQUE_BATCH_ITEM_ID                   Tag = 0x420093
	UNIQUE_IDENTIFIER                      Tag = 0x420094
	USAGE_LIMITS                           Tag = 0x420095
	USAGE_LIMITS_COUNT                     Tag = 0x420096
	USAGE_LIMITS_TOTAL                     Tag = 0x420097
	USAGE_LIMITS_UNIT                      Tag = 0x420098
	USERNAME                               Tag = 0x420099
	VALIDITY_DATE                          Tag = 0x42009A
	VALIDITY_INDICATOR                     Tag = 0x42009B
	VENDOR_EXTENSION                       Tag = 0x42009C
	VENDOR_IDENTIFICATION                  Tag = 0x42009D
	WRAPPING_METHOD                        Tag = 0x42009E
	X                                      Tag = 0x42009F
	Y                                      Tag = 0x4200A0
	PASSWORD                               Tag = 0x4200A1
	// KMIP 1.1
	DEVICE_IDENTIFIER            Tag = 0x4200A2
	ENCODING_OPTION              Tag = 0x4200A3
	EXTENSION_INFORMATION        Tag = 0x4200A4
	EXTENSION_NAME               Tag = 0x4200A5
	EXTENSION_TAG                Tag = 0x4200A6
	EXTENSION_TYPE               Tag = 0x4200A7
	FRESH                        Tag = 0x4200A8
	MACHINE_IDENTIFIER           Tag = 0x4200A9
	MEDIA_IDENTIFIER             Tag = 0x4200AA
	NETWORK_IDENTIFIER           Tag = 0x4200AB
	OBJECT_GROUP_MEMBER          Tag = 0x4200AC
	CERTIFICATE_LENGTH           Tag = 0x4200AD
	DIGITAL_SIGNATURE_ALGORITHM  Tag = 0x4200AE
	CERTIFICATE_SERIAL_NUMBER    Tag = 0x4200AF
	DEVICE_SERIAL_NUMBER         Tag = 0x4200B0
	ISSUER_ALTERNATIVE_NAME      Tag = 0x4200B1
	ISSUER_DISTINGUISHED_NAME    Tag = 0x4200B2
	SUBJECT_ALTERNATIVE_NAME     Tag = 0x4200B3
	SUBJECT_DISTINGUISHED_NAME   Tag = 0x4200B4
	X_509_CERTIFICATE_IDENTIFIER Tag = 0x4200B5
	X_509_CERTIFICATE_ISSUER     Tag = 0x4200B6
	X_509_CERTIFICATE_SUBJECT    Tag = 0x4200B7
	// KMIP 1.2
	KEY_VALUE_LOCATION            Tag = 0x4200B8
	KEY_VALUE_LOCATION_VALUE      Tag = 0x4200B9
	KEY_VALUE_LOCATION_TYPE       Tag = 0x4200BA
	KEY_VALUE_PRESENT             Tag = 0x4200BB
	ORIGINAL_CREATION_DATE        Tag = 0x4200BC
	PGP_KEY                       Tag = 0x4200BD
	PGP_KEY_VERSION               Tag = 0x4200BE
	ALTERNATIVE_NAME              Tag = 0x4200BF
	ALTERNATIVE_NAME_VALUE        Tag = 0x4200C0
	ALTERNATIVE_NAME_TYPE         Tag = 0x4200C1
	DATA                          Tag = 0x4200C2
	SIGNATURE_DATA                Tag = 0x4200C3
	DATA_LENGTH                   Tag = 0x4200C4
	RANDOM_IV                     Tag = 0x4200C5
	MAC_DATA                      Tag = 0x4200C6
	ATTESTATION_TYPE              Tag = 0x4200C7
	NONCE                         Tag = 0x4200C8
	NONCE_ID                      Tag = 0x4200C9
	NONCE_VALUE                   Tag = 0x4200CA
	ATTESTATION_MEASUREMENT       Tag = 0x4200CB
	ATTESTATION_ASSERTION         Tag = 0x4200CC
	IV_LENGTH                     Tag = 0x4200CD
	TAG_LENGTH                    Tag = 0x4200CE
	FIXED_FIELD_LENGTH            Tag = 0x4200CF
	COUNTER_LENGTH                Tag = 0x4200D0
	INITIAL_COUNTER_VALUE         Tag = 0x4200D1
	INVOCATION_FIELD_LENGTH       Tag = 0x4200D2
	ATTESTATION_CAPABLE_INDICATOR Tag = 0x4200D3
	// KMIP 1.3
	OFFSET_ITEMS                      Tag = 0x4200D4
	LOCATED_ITEMS                     Tag = 0x4200D5
	CORRELATION_VALUE                 Tag = 0x4200D6
	INIT_INDICATOR                    Tag = 0x4200D7
	FINAL_INDICATOR                   Tag = 0x4200D8
	RNG_PARAMETERS                    Tag = 0x4200D9
	RNG_ALGORITHM                     Tag = 0x4200DA
	DRBG_ALGORITHM                    Tag = 0x4200DB
	FIPS186_VARIATION                 Tag = 0x4200DC
	PREDICTION_RESISTANCE             Tag = 0x4200DD
	RANDOM_NUMBER_GENERATOR           Tag = 0x4200DE
	VALIDATION_INFORMATION            Tag = 0x4200DF
	VALIDATION_AUTHORITY_TYPE         Tag = 0x4200E0
	VALIDATION_AUTHORITY_COUNTRY      Tag = 0x4200E1
	VALIDATION_AUTHORITY_URI          Tag = 0x4200E2
	VALIDATION_VERSION_MAJOR          Tag = 0x4200E3
	VALIDATION_VERSION_MINOR          Tag = 0x4200E4
	VALIDATION_TYPE                   Tag = 0x4200E5
	VALIDATION_LEVEL                  Tag = 0x4200E6
	VALIDATION_CERTIFICATE_IDENTIFIER Tag = 0x4200E7
	VALIDATION_CERTIFICATE_URI        Tag = 0x4200E8
	VALIDATION_VENDOR_URI             Tag = 0x4200E9
	VALIDATION_PROFILE                Tag = 0x4200EA
	PROFILE_INFORMATION               Tag = 0x4200EB
	PROFILE_NAME                      Tag = 0x4200EC
	SERVER_URI                        Tag = 0x4200ED
	SERVER_PORT                       Tag = 0x4200EE
	STREAMING_CAPABILITY              Tag = 0x4200EF
	ASYNCHRONOUS_CAPABILITY           Tag = 0x4200F0
	ATTESTATION_CAPABILITY            Tag = 0x4200F1
	UNWRAP_MODE                       Tag = 0x4200F2
	DESTROY_ACTION                    Tag = 0x4200F3
	SHREDDING_ALGORITHM               Tag = 0x4200F4
	RNG_MODE                          Tag = 0x4200F5
	CLIENT_REGISTRATION_METHOD        Tag = 0x4200F6
	CAPABILITY_INFORMATION            Tag = 0x4200F7
	// KMIP 1.4
	KEY_WRAP_TYPE                               Tag = 0x4200F8
	BATCH_UNDO_CAPABILITY                       Tag = 0x4200F9
	BATCH_CONTINUE_CAPABILITY                   Tag = 0x4200FA
	PKCS12_FRIENDLY_NAME                        Tag = 0x4200FB
	DESCRIPTION                                 Tag = 0x4200FC
	COMMENT                                     Tag = 0x4200FD
	AUTHENTICATED_ENCRYPTION_ADDITIONAL_DATATag     = 0x4200FE
	AUTHENTICATED_ENCRYPTION_TAG                Tag = 0x4200FF
	SALT_LENGTH                                 Tag = 0x420100
	MASK_GENERATOR                              Tag = 0x420101
	MASK_GENERATOR_HASHING_ALGORITHM            Tag = 0x420102
	P_SOURCE                                    Tag = 0x420103
	TRAILER_FIELD                               Tag = 0x420104
	CLIENT_CORRELATION_VALUE                    Tag = 0x420105
	SERVER_CORRELATION_VALUE                    Tag = 0x420106
	DIGESTED_DATA                               Tag = 0x420107
	CERTIFICATE_SUBJECT_CN                      Tag = 0x420108
	CERTIFICATE_SUBJECT_O                       Tag = 0x420109
	CERTIFICATE_SUBJECT_OU                      Tag = 0x42010A
	CERTIFICATE_SUBJECT_EMAIL                   Tag = 0x42010B
	CERTIFICATE_SUBJECT_C                       Tag = 0x42010C
	CERTIFICATE_SUBJECT_ST                      Tag = 0x42010D
	CERTIFICATE_SUBJECT_L                       Tag = 0x42010E
	CERTIFICATE_SUBJECT_UID                     Tag = 0x42010F
	CERTIFICATE_SUBJECT_SERIAL_NUMBER           Tag = 0x420110
	CERTIFICATE_SUBJECT_TITLE                   Tag = 0x420111
	CERTIFICATE_SUBJECT_DC                      Tag = 0x420112
	CERTIFICATE_SUBJECT_DN_QUALIFIER            Tag = 0x420113
	CERTIFICATE_ISSUER_CN                       Tag = 0x420114
	CERTIFICATE_ISSUER_O                        Tag = 0x420115
	CERTIFICATE_ISSUER_OU                       Tag = 0x420116
	CERTIFICATE_ISSUER_EMAIL                    Tag = 0x420117
	CERTIFICATE_ISSUER_C                        Tag = 0x420118
	CERTIFICATE_ISSUER_ST                       Tag = 0x420119
	CERTIFICATE_ISSUER_L                        Tag = 0x42011A
	CERTIFICATE_ISSUER_UID                      Tag = 0x42011B
	CERTIFICATE_ISSUER_SERIAL_NUMBER            Tag = 0x42011C
	CERTIFICATE_ISSUER_TITLE                    Tag = 0x42011D
	CERTIFICATE_ISSUER_DC                       Tag = 0x42011E
	CERTIFICATE_ISSUER_DN_QUALIFIER             Tag = 0x42011F
	SENSITIVE                                   Tag = 0x420120
	ALWAYS_SENSITIVE                            Tag = 0x420121
	EXTRACTABLE                                 Tag = 0x420122
	NEVER_EXTRACTABLE                           Tag = 0x420123
	REPLACE_EXISTING                            Tag = 0x420124
	// KMIP 2.0
	ATTRIBUTES                            Tag = 0x420125
	COMMON_ATTRIBUTES                     Tag = 0x420126
	PRIVATE_KEY_ATTRIBUTES                Tag = 0x420127
	PUBLIC_KEY_ATTRIBUTES                 Tag = 0x420128
	EXTENSION_ENUMERATION                 Tag = 0x420129
	EXTENSION_ATTRIBUTE                   Tag = 0x42012A
	EXTENSION_PARENT_STRUCTURE_TAG        Tag = 0x42012B
	EXTENSION_DESCRIPTION                 Tag = 0x42012C
	SERVER_NAME                           Tag = 0x42012D
	SERVER_SERIAL_NUMBER                  Tag = 0x42012E
	SERVER_VERSION                        Tag = 0x42012F
	SERVER_LOAD                           Tag = 0x420130
	PRODUCT_NAME                          Tag = 0x420131
	BUILD_LEVEL                           Tag = 0x420132
	BUILD_DATE                            Tag = 0x420133
	CLUSTER_INFO                          Tag = 0x420134
	ALTERNATE_FAILOVER_ENDPOINTS          Tag = 0x420135
	SHORT_UNIQUE_IDENTIFIER               Tag = 0x420136
	RESERVED                              Tag = 0x420137
	TAG                                   Tag = 0x420138
	CERTIFICATE_REQUEST_UNIQUE_IDENTIFIER Tag = 0x420139
	NIST_KEY_TYPE                         Tag = 0x42013A
	ATTRIBUTE_REFERENCE                   Tag = 0x42013B
	CURRENT_ATTRIBUTE                     Tag = 0x42013C
	NEW_ATTRIBUTE                         Tag = 0x42013D
	// 0x42013E is designated '(Reserved)' in KMIP 2.0
	// 0x42013F is designated '(Reserved)' in KMIP 2.0
	CERTIFICATE_REQUEST_VALUE Tag = 0x420140
	LOG_MESSAGE               Tag = 0x420141
	PROFILE_VERSION           Tag = 0x420142
	PROFILE_VERSION_MAJOR     Tag = 0x420143
	PROFILE_VERSION_MINOR     Tag = 0x420144
	PROTECTION_LEVEL          Tag = 0x420145
	PROTECTION_PERIOD         Tag = 0x420146
	QUANTUM_SAFE              Tag = 0x420147
	QUANTUM_SAFE_CAPABILITY   Tag = 0x420148
	TICKET                    Tag = 0x420149
	TICKET_TYPE               Tag = 0x42014A
	TICKET_VALUE              Tag = 0x42014B
	REQUEST_COUNT             Tag = 0x42014C
	RIGHTS                    Tag = 0x42014D
	OBJECTS                   Tag = 0x42014E
	OPERATIONS                Tag = 0x42014F
	RIGHT                     Tag = 0x420150
	ENDPOINT_ROLE             Tag = 0x420151
	DEFAULTS_INFORMATION      Tag = 0x420152
	OBJECT_DEFAULTS           Tag = 0x420153
	EPHEMERAL                 Tag = 0x420154
	SERVER_HASHED_PASSWORD    Tag = 0x420155
	ONE_TIME_PASSWORD         Tag = 0x420156
	HASHED_PASSWORD           Tag = 0x420157
	ADJUSTMENT_TYPE           Tag = 0x420158
	PKCS11_INTERFACE          Tag = 0x420159
	PKCS11_FUNCTION           Tag = 0x42015A
	PKCS11_INPUT_PARAMETERS   Tag = 0x42015B
	PKCS11_OUTPUT_PARAMETERS  Tag = 0x42015C
	PKCS11_RETURN_CODE        Tag = 0x42015D
	PROTECTION_STORAGE_MASK   Tag = 0x42015E
	PROTECTION_STORAGE_MASKS  Tag = 0x42015F
	INTEROP_FUNCTION          Tag = 0x420160
	INTEROP_IDENTIFIER        Tag = 0x420161
	ADJUSTMENT_VALUE          Tag = 0x420162
)

const (
	STRUCTURE    Type = 0x01
	INTEGER      Type = 0x02
	LONG_INTEGER Type = 0x03
	BIG_INTEGER  Type = 0x04
	ENUMERATION  Type = 0x05
	BOOLEAN      Type = 0x06
	TEXT_STRING  Type = 0x07
	BYTE_STRING  Type = 0x08
	DATE_TIME    Type = 0x09
	INTERVAL     Type = 0x0A
)