# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2021_05_06_021417) do

  create_table "sparkles", force: :cascade do |t|
    t.string "reason"
    t.integer "sparkler_id", null: false
    t.integer "sparklee_id", null: false
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["sparklee_id"], name: "index_sparkles_on_sparklee_id"
    t.index ["sparkler_id"], name: "index_sparkles_on_sparkler_id"
  end

  create_table "users", force: :cascade do |t|
    t.string "handle", null: false
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["handle"], name: "index_users_on_handle", unique: true
  end

end
