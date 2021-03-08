import {createStore, applyMiddleware} from 'redux'
import {composeWithDevTools} from 'redux-devtools-extension'
import thunk from 'redux-thunk'
import rootReducer from './reducer'

const store = (initialState?: any) => {
    initialState =
        JSON.parse(window.localStorage.getItem('state') || '') || initialState
    const middleware = [thunk]

    const store = createStore(
        rootReducer,
        initialState,
        composeWithDevTools(
            applyMiddleware(...middleware)
        )
    )

    //TODO: Do something with this
    // store.subscribe(() => {
    //     const state = store.getState()
    //save the config to localstorage
    //     for (const setting in state.settings.settings) {
    //        if (state.settings.settings.hasOwnProperty(setting)) {
    //            const element = state.settings.settings[setting]
    //            localStorage.setItem(element.key, element.value)
    //        }
    //    }

    //    for (const translation in state.languages.language) {
    //        if (state.languages.language.hasOwnProperty(translation)) {
    //            const trans = state.languages.language[translation]
    //          localStorage.setItem(translation, trans)
    //        }
    //   }

    //  const persist = {
    //   }
    //   localStorage.setItem('state', JSON.stringify(persist))

    // const user = {
    //     user: state.user
    // };
    // window.localStorage.setItem("user", JSON.stringify(user.user));
    //})

    return store
}

export default store