import {createStore} from 'vuex'

const store = createStore({
    state () {
        return {
            token: '',

            // 动态导航
            navigation: {
                title: '',
                path: '',
                afterName: ''
            },

            // 动态步骤条激活状态
            stepsActive: 0,

            // 结果页标题状态
            pageTitle: ''
        }
    },
    mutations: {
        SET_TOKEN: (state, token) => {
            state.token = token
            localStorage.setItem("token", token)
        },
        resetState: (state) => {
            state.token = ''
        },
        addNav(state, item) {
                state.navigation.title = item.title;
                state.navigation.path = item.path;
                state.navigation.afterName = item.afterName;
        },
        setActive(state, active) {
            state.stepsActive = active
        },
        setPageTitle(state, title) {
            state.pageTitle = title
        }
    },
    actions: {},
    modules: {}
})

export default store;