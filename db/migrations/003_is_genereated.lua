new_format = {
    {name = 'id', type = 'string'},
    {name = 'author_id', type = 'string'},
    {name = 'audio_url', type = 'string'},
    {name = 'icon_url', type = 'string'},
    {name = 'file_name', type = 'string'},
    {name = 'created_at', type = 'string'},
    {name = 'title', type = 'string'},
    {name = 'musical_instrument', type = 'string'},
    {name = 'genre', type = 'string'},
    {name = 'mood', type = 'string'},
    {name = 'tonality', type = 'string'},
    {name = 'tempo', type = 'string'},
    {name = 'style', type = 'string'},
    {name = 'is_generated', type = 'boolean', is_nullable = true }
}

box.space.sounds:format(new_format)

local my_table = box.space.sounds
local records = my_table:select()

for _, record in ipairs(records) do
    my_table.index.primary:update(record.id, {{'=', 14, false}})
end


box.space.sounds:create_index('is_generated', {
    parts = { { 'is_generated' } },
    unique = false,
    if_not_exists = false
})
