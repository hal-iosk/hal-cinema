<template>
    <section>
        <tab></tab>

        <div class="flex-container">
            <div class="container">
                <p class="container-title">上映予定作品</p>
                <ul class="scheduled-film-container">
                    <li>
                        <div class="movie">
                            <div class="movie-title-container">
                                <p class="movie-title"><span class="title">DESTINY 鎌倉ものがたり</span></p>
                                <p class="release-date">12月9日公開</p>
                            </div>
                            <div class="contents">
                                <div class="movie-image">
                                    <img src="movie1.jpg" alt="">
                                </div>
                                <p class="description">
                                    西岸良平による人気漫画「鎌倉ものがたり」を、西岸が原作者である『ALWAYS』シリーズなどの山崎貴監督が実写映画化。
                                </p>
                            </div>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </section>
</template>

<script>
import Tab from "./tab.vue"
import MovieHttp from"../../services/movie"

export default {
    name: "comingsoon",
    components:{
        Tab
    },
    mounted(){
        MovieHttp.GetMovies
        .then((res)=>{
            console.log(res)
        })
    }
}
</script>

<style>
    body,
    html {
        width: 100%;
        height: 100%;
    }
    *,
    *:before,
    *:after {
        padding: 0;
        margin: 0;
        box-sizing: border-box;
    }

    .flex-container {
        display: flex;
        justify-content: center;
    }
    .container {
        padding-top: 2rem;
    }
    .container-title {
        padding: 0.5rem;
        border-bottom: 1px solid #DDD;
    }

    .scheduled-film-container li {
        list-style: none;
    }

    .movie {
        padding: 1rem;
    }
    .movie-title-container {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0.5rem 0;
        border-bottom: 1px dashed rgb(219, 1, 1);
    }
    .movie-title-container .movie-title {
        border-left: 10px solid rgb(219, 1, 1);
        font-size: 1.5rem;
    }
    .movie-title-container .movie-title .title {
        padding-left: 1rem;
    }
    .movie-title-container .release-date {
        font-size: 0.9rem;
        color: #555;
    }
    .contents {
        display: flex;
        justify-content: space-between;
        padding: 0.5rem;
    }
    .contents .movie-image {
        width: 150px;
    }
    .contents .movie-image img {
        display: block;
        width: 100%;
    }
    .contents .description {
        max-width: 600px;
        padding-left: 2rem;
        line-height: 1.5rem;
        letter-spacing: 0.1rem;
    }
</style>