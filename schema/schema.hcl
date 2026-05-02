table "samples" {
  schema = schema.public
  column "id" {
    null    = false
    type    = uuid
    default = sql("gen_random_uuid()")
  }
  column "message" {
    null    = false
    type    = text
    default = ""
  }
  column "created_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updated_at" {
    null    = false
    type    = timestamptz(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
}
function "set_updated_at" {
  schema = schema.public
  lang   = PLpgSQL
  return = trigger
  as     = <<-SQL
  BEGIN
    IF (TG_OP = 'UPDATE') THEN
      NEW.updated_at := CURRENT_TIMESTAMP;
      return NEW;
    END IF;
  END;
  SQL
}
trigger "set_updated_at" {
  on = table.samples
  before {
    update = true
  }
  for = ROW
  execute {
    function = function.set_updated_at
  }
}
schema "public" {
  comment = "standard public schema"
}
