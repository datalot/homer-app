import LivingBeing from './living_being';

const table = 'hep';

/**
 * A class to handle users in DB
 */
class SearchData extends LivingBeing {
  /**
   * Class constructor
   *
   * @param {object} server of hapi
   * @param {object} param of search
   */
  constructor(server, param) {
    super({db: server.databases.data, table, param});
    this.param = 1;
    this.dataDb = server.databases.data;
  }

  /*
  select * from hep where (hep_header->>'payloadType')::int = 1 limit 1;
  return knex('books').select(knex.raw("data->'author' as author"))
    .whereRaw("data->'author'->>'first_name'=? ",[books[0].author.first_name])
  */
  get(columns) {
    return this.dataDb(table)
      .whereRaw('(hep_header->>"payloadType")::int = ? ', this.param)
      .select(columns)
      .then(function(rows) {
        console.log('Found !');
        rows.forEach(function(row) {
          console.log(row);
        });
        return rows;
      });
  }
}

export default SearchData;
