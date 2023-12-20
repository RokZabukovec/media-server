import {defineStore} from 'pinia'
import axios from 'axios'

export const useMovieStore = defineStore('movieStore', {
    state: () => ({
        movies: []
    }),
    actions: {
        async fetchMovies() {
            try {
                this.movies = await axios.get('/api/files')
            } catch (error) {
                console.error('Error fetching items:', error)
            }
        }
    }
})
