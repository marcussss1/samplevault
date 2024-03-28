box.schema.space.create('sounds', {
    if_not_exists = true
})

box.space.sounds:format({
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
})

box.space.sounds:create_index('primary', {
    type = 'tree',
    parts = {'id'},
    if_not_exists = true
})

box.space.sounds:create_index('author_id', {
  parts = { { 'author_id' } },
  unique = false,
  if_not_exists = true
})

box.space.sounds:create_index('file_name', {
    parts = { { 'file_name' } },
    unique = true,
    if_not_exists = true
})

box.space.sounds:create_index('is_generated', {
    parts = { { 'is_generated' } },
    unique = false,
    if_not_exists = false
})

----------------------------------------------------------------------------------------------------

box.schema.space.create('playlists', {
    if_not_exists = true
})

box.space.playlists:format({
    {name = 'id', type = 'string'},
    {name = 'author_id', type = 'string'},
    {name = 'icon_url', type = 'string'},
    {name = 'created_at', type = 'string'},
    {name = 'title', type = 'string'},
    {name = 'description', type = 'string'},
    {name = 'sound_ids', type = 'array'}
})

box.space.playlists:create_index('primary', {
    type = 'tree',
    parts = {'id'},
    if_not_exists = true
})


--box.space.samples:create_index('author_id', {
--    type = 'tree',
--    parts = {'author_id'},
--    unique = false,
--    if_not_exists = true
--})

--box.schema.space.create('kits', {
--    engine = 'vinyl',
--    if_not_exists = true
--})
--
--box.space.kits:format({
--    {name = 'id', type = 'string'},
--    {name = 'icon_url', type = 'string'},
--    {name = 'title', type = 'string'},
--    {name = 'description', type = 'string'},
--    {name = 'created_at', type = 'string'},
--    {name = 'sample_ids', type = 'array'}
--})
--
--box.space.kits:create_index('primary', {
--    type = 'tree',
--    parts = {'id'},
--    if_not_exists = true
--})

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
