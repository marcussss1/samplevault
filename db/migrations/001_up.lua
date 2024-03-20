box.schema.space.create('sounds', {
    engine = 'vinyl',
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
    {name = 'style', type = 'string'}
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

----------------------------------------------------------------------------------------------------

box.schema.space.create('playlists', {
    engine = 'vinyl',
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

box.space.playlists:create_index('author_id', {
    type = 'tree',
    parts = { { 'author_id' } },
    unique = false,
    if_not_exists = true
})


box.space.playlists:update('playlist_id_123', {{'=', 7, {'ef383bf8-eee3-4923-801c-b3daf4397676','ddc4c240-145b-4a5c-b324-c61abc1465aa'}}})

box.space.playlists:insert{
    'playlist_id_456',         -- id
    '0332a44c2d178be5927cb2615a915ecb2fd062d0ca17aade7a5bd09d3badc80a',           -- author_id
    'https://example.com/icon.jpg',  -- icon_url
    '2022-01-01 12:00:00',     -- created_at
    'My Playlist 2',             -- title
    'Description of the playlist 2',  -- description
    {'ef383bf8-eee3-4923-801c-b3daf4397676'}                  -- sound_ids (array)
}


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
