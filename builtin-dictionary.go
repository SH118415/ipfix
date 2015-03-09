package ipfix

// Autogenerated Mon Mar  9 20:50:08 CET 2015
var builtinDictionary = fieldDictionary{
	dictionaryKey{0, 141}: DictionaryEntry{FieldID: 141, Name: "lineCardId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 142}: DictionaryEntry{FieldID: 142, Name: "portId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 10}:  DictionaryEntry{FieldID: 10, Name: "ingressInterface", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 14}:  DictionaryEntry{FieldID: 14, Name: "egressInterface", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 143}: DictionaryEntry{FieldID: 143, Name: "meteringProcessId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 144}: DictionaryEntry{FieldID: 144, Name: "exportingProcessId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 148}: DictionaryEntry{FieldID: 148, Name: "flowId", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 145}: DictionaryEntry{FieldID: 145, Name: "templateId", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 149}: DictionaryEntry{FieldID: 149, Name: "observationDomainId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 138}: DictionaryEntry{FieldID: 138, Name: "observationPointId", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 137}: DictionaryEntry{FieldID: 137, Name: "commonPropertiesId", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 130}: DictionaryEntry{FieldID: 130, Name: "exporterIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 131}: DictionaryEntry{FieldID: 131, Name: "exporterIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 217}: DictionaryEntry{FieldID: 217, Name: "exporterTransportPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 211}: DictionaryEntry{FieldID: 211, Name: "collectorIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 212}: DictionaryEntry{FieldID: 212, Name: "collectorIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 213}: DictionaryEntry{FieldID: 213, Name: "exportInterface", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 214}: DictionaryEntry{FieldID: 214, Name: "exportProtocolVersion", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 215}: DictionaryEntry{FieldID: 215, Name: "exportTransportProtocol", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 216}: DictionaryEntry{FieldID: 216, Name: "collectorTransportPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 173}: DictionaryEntry{FieldID: 173, Name: "flowKeyIndicator", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 41}:  DictionaryEntry{FieldID: 41, Name: "exportedMessageTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 40}:  DictionaryEntry{FieldID: 40, Name: "exportedOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 42}:  DictionaryEntry{FieldID: 42, Name: "exportedFlowRecordTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 163}: DictionaryEntry{FieldID: 163, Name: "observedFlowTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 164}: DictionaryEntry{FieldID: 164, Name: "ignoredPacketTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 165}: DictionaryEntry{FieldID: 165, Name: "ignoredOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 166}: DictionaryEntry{FieldID: 166, Name: "notSentFlowTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 167}: DictionaryEntry{FieldID: 167, Name: "notSentPacketTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 168}: DictionaryEntry{FieldID: 168, Name: "notSentOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 60}:  DictionaryEntry{FieldID: 60, Name: "ipVersion", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 8}:   DictionaryEntry{FieldID: 8, Name: "sourceIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 27}:  DictionaryEntry{FieldID: 27, Name: "sourceIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 9}:   DictionaryEntry{FieldID: 9, Name: "sourceIPv4PrefixLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 29}:  DictionaryEntry{FieldID: 29, Name: "sourceIPv6PrefixLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 44}:  DictionaryEntry{FieldID: 44, Name: "sourceIPv4Prefix", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 170}: DictionaryEntry{FieldID: 170, Name: "sourceIPv6Prefix", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 12}:  DictionaryEntry{FieldID: 12, Name: "destinationIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 28}:  DictionaryEntry{FieldID: 28, Name: "destinationIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 13}:  DictionaryEntry{FieldID: 13, Name: "destinationIPv4PrefixLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 30}:  DictionaryEntry{FieldID: 30, Name: "destinationIPv6PrefixLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 45}:  DictionaryEntry{FieldID: 45, Name: "destinationIPv4Prefix", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 169}: DictionaryEntry{FieldID: 169, Name: "destinationIPv6Prefix", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 192}: DictionaryEntry{FieldID: 192, Name: "ipTTL", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 4}:   DictionaryEntry{FieldID: 4, Name: "protocolIdentifier", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 193}: DictionaryEntry{FieldID: 193, Name: "nextHeaderIPv6", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 195}: DictionaryEntry{FieldID: 195, Name: "ipDiffServCodePoint", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 196}: DictionaryEntry{FieldID: 196, Name: "ipPrecedence", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 5}:   DictionaryEntry{FieldID: 5, Name: "ipClassOfService", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 55}:  DictionaryEntry{FieldID: 55, Name: "postIpClassOfService", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 31}:  DictionaryEntry{FieldID: 31, Name: "flowLabelIPv6", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 206}: DictionaryEntry{FieldID: 206, Name: "isMulticast", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 54}:  DictionaryEntry{FieldID: 54, Name: "fragmentIdentification", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 88}:  DictionaryEntry{FieldID: 88, Name: "fragmentOffset", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 197}: DictionaryEntry{FieldID: 197, Name: "fragmentFlags", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 189}: DictionaryEntry{FieldID: 189, Name: "ipHeaderLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 207}: DictionaryEntry{FieldID: 207, Name: "ipv4IHL", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 190}: DictionaryEntry{FieldID: 190, Name: "totalLengthIPv4", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 224}: DictionaryEntry{FieldID: 224, Name: "ipTotalLength", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 191}: DictionaryEntry{FieldID: 191, Name: "payloadLengthIPv6", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 7}:   DictionaryEntry{FieldID: 7, Name: "sourceTransportPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 11}:  DictionaryEntry{FieldID: 11, Name: "destinationTransportPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 180}: DictionaryEntry{FieldID: 180, Name: "udpSourcePort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 181}: DictionaryEntry{FieldID: 181, Name: "udpDestinationPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 205}: DictionaryEntry{FieldID: 205, Name: "udpMessageLength", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 182}: DictionaryEntry{FieldID: 182, Name: "tcpSourcePort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 183}: DictionaryEntry{FieldID: 183, Name: "tcpDestinationPort", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 184}: DictionaryEntry{FieldID: 184, Name: "tcpSequenceNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 185}: DictionaryEntry{FieldID: 185, Name: "tcpAcknowledgementNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 186}: DictionaryEntry{FieldID: 186, Name: "tcpWindowSize", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 238}: DictionaryEntry{FieldID: 238, Name: "tcpWindowScale", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 187}: DictionaryEntry{FieldID: 187, Name: "tcpUrgentPointer", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 188}: DictionaryEntry{FieldID: 188, Name: "tcpHeaderLength", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 32}:  DictionaryEntry{FieldID: 32, Name: "icmpTypeCodeIPv4", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 176}: DictionaryEntry{FieldID: 176, Name: "icmpTypeIPv4", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 177}: DictionaryEntry{FieldID: 177, Name: "icmpCodeIPv4", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 139}: DictionaryEntry{FieldID: 139, Name: "icmpTypeCodeIPv6", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 178}: DictionaryEntry{FieldID: 178, Name: "icmpTypeIPv6", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 179}: DictionaryEntry{FieldID: 179, Name: "icmpCodeIPv6", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 33}:  DictionaryEntry{FieldID: 33, Name: "igmpType", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 56}:  DictionaryEntry{FieldID: 56, Name: "sourceMacAddress", Type: FieldTypes["macAddress"]},
	dictionaryKey{0, 81}:  DictionaryEntry{FieldID: 81, Name: "postSourceMacAddress", Type: FieldTypes["macAddress"]},
	dictionaryKey{0, 58}:  DictionaryEntry{FieldID: 58, Name: "vlanId", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 59}:  DictionaryEntry{FieldID: 59, Name: "postVlanId", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 80}:  DictionaryEntry{FieldID: 80, Name: "destinationMacAddress", Type: FieldTypes["macAddress"]},
	dictionaryKey{0, 57}:  DictionaryEntry{FieldID: 57, Name: "postDestinationMacAddress", Type: FieldTypes["macAddress"]},
	dictionaryKey{0, 146}: DictionaryEntry{FieldID: 146, Name: "wlanChannelId", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 147}: DictionaryEntry{FieldID: 147, Name: "wlanSSID", Type: FieldTypes["string"]},
	dictionaryKey{0, 200}: DictionaryEntry{FieldID: 200, Name: "mplsTopLabelTTL", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 203}: DictionaryEntry{FieldID: 203, Name: "mplsTopLabelExp", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 237}: DictionaryEntry{FieldID: 237, Name: "postMplsTopLabelExp", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 202}: DictionaryEntry{FieldID: 202, Name: "mplsLabelStackDepth", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 201}: DictionaryEntry{FieldID: 201, Name: "mplsLabelStackLength", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 194}: DictionaryEntry{FieldID: 194, Name: "mplsPayloadLength", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 70}:  DictionaryEntry{FieldID: 70, Name: "mplsTopLabelStackSection", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 71}:  DictionaryEntry{FieldID: 71, Name: "mplsLabelStackSection2", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 72}:  DictionaryEntry{FieldID: 72, Name: "mplsLabelStackSection3", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 73}:  DictionaryEntry{FieldID: 73, Name: "mplsLabelStackSection4", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 74}:  DictionaryEntry{FieldID: 74, Name: "mplsLabelStackSection5", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 75}:  DictionaryEntry{FieldID: 75, Name: "mplsLabelStackSection6", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 76}:  DictionaryEntry{FieldID: 76, Name: "mplsLabelStackSection7", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 77}:  DictionaryEntry{FieldID: 77, Name: "mplsLabelStackSection8", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 78}:  DictionaryEntry{FieldID: 78, Name: "mplsLabelStackSection9", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 79}:  DictionaryEntry{FieldID: 79, Name: "mplsLabelStackSection10", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 204}: DictionaryEntry{FieldID: 204, Name: "ipPayloadLength", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 15}:  DictionaryEntry{FieldID: 15, Name: "ipNextHopIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 62}:  DictionaryEntry{FieldID: 62, Name: "ipNextHopIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 16}:  DictionaryEntry{FieldID: 16, Name: "bgpSourceAsNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 17}:  DictionaryEntry{FieldID: 17, Name: "bgpDestinationAsNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 128}: DictionaryEntry{FieldID: 128, Name: "bgpNextAdjacentAsNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 129}: DictionaryEntry{FieldID: 129, Name: "bgpPrevAdjacentAsNumber", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 18}:  DictionaryEntry{FieldID: 18, Name: "bgpNextHopIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 63}:  DictionaryEntry{FieldID: 63, Name: "bgpNextHopIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 46}:  DictionaryEntry{FieldID: 46, Name: "mplsTopLabelType", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 47}:  DictionaryEntry{FieldID: 47, Name: "mplsTopLabelIPv4Address", Type: FieldTypes["ipv4Address"]},
	dictionaryKey{0, 140}: DictionaryEntry{FieldID: 140, Name: "mplsTopLabelIPv6Address", Type: FieldTypes["ipv6Address"]},
	dictionaryKey{0, 90}:  DictionaryEntry{FieldID: 90, Name: "mplsVpnRouteDistinguisher", Type: FieldTypes["octetArray"]},
	dictionaryKey{0, 25}:  DictionaryEntry{FieldID: 25, Name: "minimumIpTotalLength", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 26}:  DictionaryEntry{FieldID: 26, Name: "maximumIpTotalLength", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 52}:  DictionaryEntry{FieldID: 52, Name: "minimumTTL", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 53}:  DictionaryEntry{FieldID: 53, Name: "maximumTTL", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 208}: DictionaryEntry{FieldID: 208, Name: "ipv4Options", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 64}:  DictionaryEntry{FieldID: 64, Name: "ipv6ExtensionHeaders", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 6}:   DictionaryEntry{FieldID: 6, Name: "tcpControlBits", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 209}: DictionaryEntry{FieldID: 209, Name: "tcpOptions", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 150}: DictionaryEntry{FieldID: 150, Name: "flowStartSeconds", Type: FieldTypes["dateTimeSeconds"]},
	dictionaryKey{0, 151}: DictionaryEntry{FieldID: 151, Name: "flowEndSeconds", Type: FieldTypes["dateTimeSeconds"]},
	dictionaryKey{0, 152}: DictionaryEntry{FieldID: 152, Name: "flowStartMilliseconds", Type: FieldTypes["dateTimeMilliseconds"]},
	dictionaryKey{0, 153}: DictionaryEntry{FieldID: 153, Name: "flowEndMilliseconds", Type: FieldTypes["dateTimeMilliseconds"]},
	dictionaryKey{0, 154}: DictionaryEntry{FieldID: 154, Name: "flowStartMicroseconds", Type: FieldTypes["dateTimeMicroseconds"]},
	dictionaryKey{0, 155}: DictionaryEntry{FieldID: 155, Name: "flowEndMicroseconds", Type: FieldTypes["dateTimeMicroseconds"]},
	dictionaryKey{0, 156}: DictionaryEntry{FieldID: 156, Name: "flowStartNanoseconds", Type: FieldTypes["dateTimeNanoseconds"]},
	dictionaryKey{0, 157}: DictionaryEntry{FieldID: 157, Name: "flowEndNanoseconds", Type: FieldTypes["dateTimeNanoseconds"]},
	dictionaryKey{0, 158}: DictionaryEntry{FieldID: 158, Name: "flowStartDeltaMicroseconds", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 159}: DictionaryEntry{FieldID: 159, Name: "flowEndDeltaMicroseconds", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 160}: DictionaryEntry{FieldID: 160, Name: "systemInitTimeMilliseconds", Type: FieldTypes["dateTimeMilliseconds"]},
	dictionaryKey{0, 22}:  DictionaryEntry{FieldID: 22, Name: "flowStartSysUpTime", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 21}:  DictionaryEntry{FieldID: 21, Name: "flowEndSysUpTime", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 1}:   DictionaryEntry{FieldID: 1, Name: "octetDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 23}:  DictionaryEntry{FieldID: 23, Name: "postOctetDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 198}: DictionaryEntry{FieldID: 198, Name: "octetDeltaSumOfSquares", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 85}:  DictionaryEntry{FieldID: 85, Name: "octetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 171}: DictionaryEntry{FieldID: 171, Name: "postOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 199}: DictionaryEntry{FieldID: 199, Name: "octetTotalSumOfSquares", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 2}:   DictionaryEntry{FieldID: 2, Name: "packetDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 24}:  DictionaryEntry{FieldID: 24, Name: "postPacketDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 86}:  DictionaryEntry{FieldID: 86, Name: "packetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 172}: DictionaryEntry{FieldID: 172, Name: "postPacketTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 132}: DictionaryEntry{FieldID: 132, Name: "droppedOctetDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 133}: DictionaryEntry{FieldID: 133, Name: "droppedPacketDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 134}: DictionaryEntry{FieldID: 134, Name: "droppedOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 135}: DictionaryEntry{FieldID: 135, Name: "droppedPacketTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 19}:  DictionaryEntry{FieldID: 19, Name: "postMCastPacketDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 20}:  DictionaryEntry{FieldID: 20, Name: "postMCastOctetDeltaCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 174}: DictionaryEntry{FieldID: 174, Name: "postMCastPacketTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 175}: DictionaryEntry{FieldID: 175, Name: "postMCastOctetTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 218}: DictionaryEntry{FieldID: 218, Name: "tcpSynTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 219}: DictionaryEntry{FieldID: 219, Name: "tcpFinTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 220}: DictionaryEntry{FieldID: 220, Name: "tcpRstTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 221}: DictionaryEntry{FieldID: 221, Name: "tcpPshTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 222}: DictionaryEntry{FieldID: 222, Name: "tcpAckTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 223}: DictionaryEntry{FieldID: 223, Name: "tcpUrgTotalCount", Type: FieldTypes["unsigned64"]},
	dictionaryKey{0, 36}:  DictionaryEntry{FieldID: 36, Name: "flowActiveTimeout", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 37}:  DictionaryEntry{FieldID: 37, Name: "flowIdleTimeout", Type: FieldTypes["unsigned16"]},
	dictionaryKey{0, 136}: DictionaryEntry{FieldID: 136, Name: "flowEndReason", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 161}: DictionaryEntry{FieldID: 161, Name: "flowDurationMilliseconds", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 162}: DictionaryEntry{FieldID: 162, Name: "flowDurationMicroseconds", Type: FieldTypes["unsigned32"]},
	dictionaryKey{0, 61}:  DictionaryEntry{FieldID: 61, Name: "flowDirection", Type: FieldTypes["unsigned8"]},
	dictionaryKey{0, 210}: DictionaryEntry{FieldID: 210, Name: "paddingOctets", Type: FieldTypes["octetArray"]},
}
