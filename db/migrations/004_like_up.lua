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
    {name = 'is_generated', type = 'boolean'},
    {name = 'likes', type = 'integer', is_nullable = true}
}

box.space.sounds:format(new_format)

--local records = box.space.sounds:select()
--
--for _, record in ipairs(records) do
--    box.space.sounds:update(record[1], {
--        {'=', 'likes', 0},
--    })
--end

----------------------------------------------------------------------------------------------------

box.schema.space.create('likes', {
    if_not_exists = true
})

box.space.likes:format({
    {name = 'id', type = 'string'},
    {name = 'author_id', type = 'string'},
    {name = 'sound_id', type = 'string'},
})

box.space.likes:create_index('primary', {
    --type = 'tree',
    parts = {'id'},
    if_not_exists = true
})

box.space.likes:create_index('author_id', {
    --type = 'tree',
    parts = {'author_id'},
    if_not_exists = true,
    unique = false
})

box.space.likes:create_index('sound_id', {
    --type = 'tree',
    parts = {'sound_id'},
    if_not_exists = true,
    unique = false
})
