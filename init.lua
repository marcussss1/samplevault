local user = os.getenv("TARANTOOL_USER")
local password = os.getenv("TARANTOOL_PASSWORD")
local database = os.getenv("TARANTOOL_DATABASE")

box.cfg {
    listen = 3301,
    vinyl_dir = '/var/lib/tarantool',
    wal_mode = 'write',
    checkpoint_count = 5,
    checkpoint_interval = 48000,
}

box.schema.user.create(user, { password = password, if_not_exists = true })
box.schema.create_space(database, { if_not_exists = true })
box.schema.user.grant(user, 'read,write,execute,create,drop', 'universe', nil, { if_not_exists = true })
