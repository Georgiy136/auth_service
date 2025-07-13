box.cfg{}

-- Создание пространства для хранения сессий
local sessions_space = box.schema.space.create('sessions', {if_not_exists = true})
sessions_space:create_index('primary', {
    type = 'hash',
    parts = {'session_id'},
    if_not_exists = true
})

-- Функция для сохранения сессии
function save_session(user_id, session_id, user_agent, ip_address)
    local session_tuple = {user_id, session_id, user_agent, ip_address}
    sessions_space:insert(session_tuple)
    return true
end

-- Функция для получения сессии по session_id
function get_session(session_id)
    return sessions_space:select(session_id)
end