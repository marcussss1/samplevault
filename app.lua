local user = os.getenv("TARANTOOL_USER")
local password = os.getenv("TARANTOOL_PASSWORD")
local database = os.getenv("TARANTOOL_DATABASE")

box.cfg {
    listen = 3301,
    vinyl_dir = '/var/lib/tarantool',
    wal_mode = 'write',
    checkpoint_count = 5,
    checkpoint_interval = 60,
}

box.schema.user.create(user, { password = password, if_not_exists = true })
box.schema.create_space(database, { if_not_exists = true })
box.schema.user.grant(user, 'read,write,execute,create,drop', 'universe', nil, { if_not_exists = true })

-- Print confirmation message
print('Database initialized')

--print('Ready 1')
--
--box.schema.space.create('bands', { engine = 'vinyl' })
--
--print('Ready 2')
--
--box.space.bands:format({
--    { name = 'id', type = 'unsigned' },
--    { name = 'band_name', type = 'string' },
--    { name = 'year', type = 'unsigned' }
--})
--
--print('Ready 3')
--
--box.space.bands:create_index('primary', { type = "tree", parts = { 'id' } })
--
--print('Ready 4')
--
--box.space.bands:insert { 1, 'Roxette', 1986 }
--box.space.bands:insert { 2, 'Scorpions', 1965 }
--box.space.bands:insert { 3, 'Ace of Base', 1987 }
--
--print('Ready 5')
--
--box.space.bands:select { 3 }
