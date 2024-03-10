box.cfg {
    vinyl_dir = '/var/lib/tarantool',
}

box.schema.space.create('my_space', {if_not_exists = true})
box.space.my_space:create_index('primary', {type = 'tree', parts = {1, 'unsigned'}})
