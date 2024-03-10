box.cfg {
    listen = 3301,
    vinyl_dir = '/tmp/tarantool',
}

box.schema.space.create('bands')

box.space.bands:format({
    { name = 'id', type = 'unsigned' },
    { name = 'band_name', type = 'string' },
    { name = 'year', type = 'unsigned' }
})

box.space.bands:create_index('primary', { type = "tree", parts = { 'id' } })
