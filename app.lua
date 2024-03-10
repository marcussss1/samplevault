box.cfg {
    listen = 3301,
    vinyl_dir = '/var/lib/tarantool',
    wal_mode = 'write',
    checkpoint_count = 5,
    checkpoint_interval = 3600,
}

box.schema.space.create('bands', { engine = 'vinyl' })

box.space.bands:format({
    { name = 'id', type = 'unsigned' },
    { name = 'band_name', type = 'string' },
    { name = 'year', type = 'unsigned' }
})

box.space.bands:create_index('primary', { type = "tree", parts = { 'id' } })

box.space.bands:insert { 1, 'Roxette', 1986 }
box.space.bands:insert { 2, 'Scorpions', 1965 }
box.space.bands:insert { 3, 'Ace of Base', 1987 }
