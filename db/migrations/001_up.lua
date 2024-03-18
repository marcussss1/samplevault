box.schema.space.create('samples', { engine = 'vinyl', if_not_exists = true })
box.space.samples:format({
    {name = 'id', type = 'string'},
    {name = 'author_id', type = 'string'},
    {name = 'audio_url', type = 'string'},
    {name = 'icon_url', type = 'string'},
    {name = 'file_name', type = 'string'},
    {name = 'title', type = 'string'},
    {name = 'duration', type = 'string'},
    {name = 'musical_instrument', type = 'string'},
    {name = 'genre', type = 'string'},
    {name = 'is_favourite', type = 'boolean'},
    {name = 'created_at', type = 'string'}
})

box.space.samples:create_index('primary', {
    type = 'tree',
    parts = {'author_id'},
    if_not_exists = true
})

--box.space.samples:insert{
--    'a2802d62-b006-4949-8fa0-07328bd26719',
--    'a2802d62-b006-4949-8fa0-07328bd26719',
--    'audio_url',
--    'icon_url',
--    'Название сэмпла',
--    'Длительность сэмпла',
--    'Музыкальный инструмент сэмпла',
--    'Жанр сэмпла',
--    true,
--}
