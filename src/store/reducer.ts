import {combineReducers} from "redux";
import uploadsReducer from '../services/uploads/reducer'

export default combineReducers({
    uploads: uploadsReducer
})
// const initialState: any = {
//     articles: [
//         {
//             id: 1,
//             title: "post 1",
//             body:
//                 "Quisque cursus, metus vitae pharetra Nam libero tempore, cum soluta nobis est eligendi",
//         },
//         {
//             id: 2,
//             title: "post 2",
//             body:
//                 "Harum quidem rerum facilis est et expedita distinctio quas molestias excepturi sint",
//         },
//     ],
// }
//
// const reducer = (
//     state: any = initialState,
//     action: any
// ): any => {
//     switch (action.type) {
//         case actionTypes.ADD_ARTICLE:
//             const newArticle: any = {
//                 id: Math.random(), // not really unique
//                 title: action.article.title,
//                 body: action.article.body,
//             }
//             return {
//                 ...state,
//                 articles: state.articles.concat(newArticle),
//             }
//         case actionTypes.REMOVE_ARTICLE:
//             const updatedArticles: any[] = state.articles.filter(
//                 (article: any) => article.id !== action.article.id
//             )
//             return {
//                 ...state,
//                 articles: updatedArticles,
//             }
//     }
//     return state
// }
//
// export default reducer