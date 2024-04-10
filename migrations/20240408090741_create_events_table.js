/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable("events", function (table) {
    table.increments("event_id").primary();
    table.string("event_name");
    table.date("start_date");
    table.date("end_date");
    table.string("location");
    table.string("description");
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTable("events");
};
