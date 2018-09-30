const uuidv4 = require('uuid/v4');

exports.seed = function seed(knex) {
  const tableName = 'mapping_schema';

  let empty = {};

  let fieldsMapping1default = [
    {
      id: 'sid',
      type: 'string',
      index: 'secondary',
      name: 'Session ID',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocolFamily',
      name: 'Proto Family',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocol',
      name: 'Protocol Type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcIp',
      name: 'Source IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstIp',
      name: 'Destination IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcPort',
      name: 'Src Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstPort',
      name: 'Dst Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeSeconds',
      name: 'Timeseconds',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeUseconds',
      name: 'Usecond time',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.payloadType',
      name: 'Payload type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.captureId',
      name: 'Capture ID',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.capturePass',
      name: 'Capture Pass',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.correlation_id',
      name: 'Correlation ID',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.method',
      name: 'SIP Method',
      type: 'string',
      index: 'none',
      form_type: 'input',
      form_default: ['INVITE', 'BYE', '100', '200', '183', 'CANCEL'],
    },
    {
      id: 'data_header.callid',
      name: 'SIP Callid',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.cseq',
      name: 'SIP Cseq',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.to_user',
      name: 'SIP To user',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.from_tag',
      name: 'SIP From tag',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.protocol',
      name: 'SIP Protocol',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.from_user',
      name: 'SIP From user',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'raw',
      name: 'SIP RAW',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
  ];
  
  let fieldsMapping100default = [
    {
      id: 'sid',
      type: 'string',
      index: 'secondary',
      name: 'Session ID',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocolFamily',
      name: 'Proto Family',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocol',
      name: 'Protocol Type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcIp',
      name: 'Source IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstIp',
      name: 'Destination IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcPort',
      name: 'Src Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstPort',
      name: 'Dst Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeSeconds',
      name: 'Timeseconds',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeUseconds',
      name: 'Usecond time',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.payloadType',
      name: 'Payload type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.captureId',
      name: 'Capture ID',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.capturePass',
      name: 'Capture Pass',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.correlation_id',
      name: 'Correlation ID',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'raw',
      name: 'RAW',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
  ];
  
  let fieldsMapping34default = [
    {
      id: 'sid',
      type: 'string',
      index: 'secondary',
      name: 'Session ID',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocolFamily',
      name: 'Proto Family',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.protocol',
      name: 'Protocol Type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcIp',
      name: 'Source IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstIp',
      name: 'Destination IP',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.srcPort',
      name: 'Src Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.dstPort',
      name: 'Dst Port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeSeconds',
      name: 'Timeseconds',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.timeUseconds',
      name: 'Usecond time',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.payloadType',
      name: 'Payload type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.captureId',
      name: 'Capture ID',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.capturePass',
      name: 'Capture Pass',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.correlation_id',
      name: 'Correlation ID',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'raw',
      name: 'RAW',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
  ];
  
  let fieldsMapping1000default = [
    {
      id: 'sid',
      type: 'string',
      index: 'secondary',
      name: 'Session ID',
      form_type: 'input',
    },
    {
      id: 'protocol_header.address',
      name: 'Proto Address',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.family',
      name: 'Family',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'protocol_header.port',
      name: 'Protocol port',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.type',
      name: 'Data type',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.handle',
      name: 'Data Handle',
      type: 'integer',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.medium',
      name: 'Data Medium',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.source',
      name: 'Data Source',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'data_header.session',
      name: 'Data Session',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
    {
      id: 'raw',
      name: 'RAW',
      type: 'string',
      index: 'none',
      form_type: 'input',
    },
  ];
  
  let correlationMapping1default = [
    {
      source_field: 'data_header.callid',
      lookup_id: 100,
      lookup_profile: 'default',
      lookup_field: 'sid',
      lookup_range: [-300, 200],
    },
    {
      source_field: 'data_header.callid',
      lookup_id: 5,
      lookup_profile: 'default',
      lookup_field: 'sid',
      lookup_range: [-300, 200],
    },
  ];
    
  let correlationMapping100default = [
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'call',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'registration',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'default',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
  ];
  
  let correlationMapping34default = [
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'call',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
  ];
  
  let correlationMapping5default = [
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'call',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
  ];
  
  let correlationMapping1000default = [
    {
      source_field: 'sid',
      lookup_id: 1,
      lookup_profile: 'call',
      lookup_field: 'data_header.callid',
      lookup_range: [-300, 200],
    },
  ];
            
  const rows = [
    {
      guid: uuidv4(),
      profile: 'default',
      hepid: 1,
      hep_alias: 'SIP',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping1default),
      correlation_mapping: JSON.stringify(correlationMapping1default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      profile: 'call',
      hepid: 1,
      hep_alias: 'SIP',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping1default),
      correlation_mapping: JSON.stringify(correlationMapping1default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      profile: 'registration',
      hepid: 1,
      hep_alias: 'SIP',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping1default),
      correlation_mapping: JSON.stringify(correlationMapping1default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      profile: 'default',
      hepid: 100,
      hep_alias: 'LOG',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping100default),
      correlation_mapping: JSON.stringify(correlationMapping100default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      profile: 'default',
      hepid: 34,
      hep_alias: 'RTP-FULL-REPORT',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping34default),
      correlation_mapping: JSON.stringify(correlationMapping34default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      profile: 'default',
      hepid: 1000,
      hep_alias: 'JANUS',
      partid: 10,
      version: 1,
      retention: 10,
      partition_step: 10,
      create_index: JSON.stringify(empty),
      create_table: 'CREATE TABLE test(id integer, data text);',
      fields_mapping: JSON.stringify(fieldsMapping1000default),
      correlation_mapping: JSON.stringify(correlationMapping1000default),
      schema_mapping: JSON.stringify(empty),
      schema_settings: JSON.stringify(empty),
      create_date: new Date(),
    },
    
  ];

  return knex(tableName)
    // Empty the table (DELETE)
    .del()
    .then(function() {
      return knex.insert(rows).into(tableName);
    });
};
