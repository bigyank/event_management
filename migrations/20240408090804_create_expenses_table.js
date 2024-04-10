/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.createTable("expenses", function (table) {
    table.increments("expense_id").primary();
    table.integer("event_id").unsigned();
    table.string("item_name");
    table.decimal("cost", 10, 2);
    table.string("description");
    table.string("category");
    table.foreign("event_id").references("event_id").inTable("events");
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.dropTable("expenses");
};
