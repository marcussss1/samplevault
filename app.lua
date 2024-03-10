box.cfg {
    listen = 3301,
    vinyl_dir = '/tmp/tarantool',
}

box.schema.space.create('samplevault', {if_not_exists = true})
box.space.my_space:create_index('primary', {type = 'tree', parts = {1, 'unsigned'}})

