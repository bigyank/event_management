/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable("event_sessions", function (table) {
    table.increments("session_id").primary();
    table.integer("event_id").unsigned();
    table.string("session_name");
    table.datetime("start_time");
    table.datetime("end_time");
    table.foreign("event_id").references("event_id").inTable("events");
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTable("event_sessions");
};
