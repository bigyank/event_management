/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable("event_organizers", function (table) {
    table.increments("event_organizer_id").primary();
    table.integer("event_id").unsigned();
    table.integer("user_id").unsigned();
    table.string("role");
    table.foreign("event_id").references("event_id").inTable("events");
    table.foreign("user_id").references("user_id").inTable("users");
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTable("event_organizers");
};
