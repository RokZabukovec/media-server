import {defineStore} from 'pinia'
import axios from 'axios'

export const useCategoryStore = defineStore('categoryStore', {
    state: () => ({
        categories: []
    }),
    actions: {
        async fetchCategories(id) {
            try {
                const url = id ? `/api/categories?id=${id}` : '/api/categories';
                this.categories = await axios.get(url);
            } catch (error) {
                console.error('Error fetching categories:', error);
            }
        }
    }
});

