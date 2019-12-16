const uuidv4 = require('uuid/v4');

exports.seed = function seed(knex) {
  const tableName = 'user_settings';

  let empty = {};
  let dashboardHome='{"id":"home","name":"Home","alias":"home","selectedItem":"","title":"Home","weight":10,"widgets":[{"x":0,"y":0,"cols":2,"rows":1,"name":"clock","title":"clock","id":"clock214","output":{},"config":{"id":"clock214","datePattern":"YYYY-MM-DD","location":{"value":-60,"offset":"+1","name":"GMT+1 CET","desc":"Central European Time"},"showseconds":false,"timePattern":"HH:mm:ss","title":"Home Clock"}},{"x":0,"y":1,"cols":2,"rows":3,"name":"display-results","title":"display-results","id":"display-results370","output":{},"config":{"id":"display-results370","title":"CALL SIP SEARCH","group":"Search","name":"protosearch","description":"Display Search Form component","refresh":false,"sizeX":2,"sizeY":2,"config":{"title":"CALL SIP SEARCH","searchbutton":true,"protocol_id":{"name":"SIP","value":1},"protocol_profile":{"name":"call","value":"call"}},"uuid":"ed426bd0-ff21-40f7-8852-58700abc3762","fields":[{"field_name":"data_header.from_user","hepid":1,"name":"1:call:data_header.from_user","selection":"SIP From user","type":"string"},{"field_name":"data_header.to_user","hepid":1,"name":"1:call:data_header.to_user","selection":"SIP To user","type":"string"},{"field_name":"data_header.method","hepid":1,"name":"1:call:data_header.method","selection":"SIP Method","type":"string"},{"field_name":"data_header.callid","hepid":1,"name":"1:call:data_header.callid","selection":"SIP Callid","type":"string"},{"field_name":"limit","hepid":1,"name":"1:call:limit","selection":"Query Limit","type":"string"},{"field_name":"targetResultsContainer","hepid":1,"name":"1:call:targetResultsContainer","selection":"Results Container","type":"string"}],"row":0,"col":1,"cols":2,"rows":2,"x":0,"y":1,"protocol_id":{"name":"SIP","value":100}}},{"x":2,"y":0,"cols":4,"rows":4,"name":"result","title":"result","id":"result560","output":{}}],"config":{"margins":[10,10],"columns":"6","pushing":true,"draggable":{"handle":".box-header"},"resizable":{"enabled":true,"handles":["n","e","s","w","ne","se","sw","nw"]}}}';
  const rows = [
    {
      guid: uuidv4(),
      username: 'admin',
      param: 'home',
      partid: 10,
      category: 'dashboard',
      data: dashboardHome,
      create_date: new Date(),
    },
    {
      guid: uuidv4(),
      username: 'support',
      param: 'home',
      partid: 10,
      category: 'dashboard',
      data: dashboardHome,
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
