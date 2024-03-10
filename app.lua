box.cfg {
    listen = 3301,
    vinyl_dir = '/var/lib/tarantool',
    wal_mode = 'none',
}

--vinyl.storage.engine = "vinyl"
--
--box.once('vinyl_checkpoint_config', function()
--    vinyl.checkpoints_path = '/var/lib/tarantool/checkpoints'
--    vinyl.checkpoint_count = 3
--end)
--
--box.cfg {
--    wal_dir = '/var/lib/tarantool/wal'
--}




--box.schema.space.create('bands', {if_not_exists = true,engine='vinyl'})
--
--box.space.bands:format({
--    { name = 'id', type = 'unsigned' },
--    { name = 'band_name', type = 'string' },
--    { name = 'year', type = 'unsigned' }
--})
--
--box.space.bands:create_index('primary', { type = "tree", parts = { 'id' } })
--
--box.space.bands:insert { 1, 'Roxette', 1986 }
--box.space.bands:insert { 2, 'Scorpions', 1965 }
--box.space.bands:insert { 3, 'Ace of Base', 1987 }
--
--box.space.bands:select { 3 }
