<template>
    <section>
        <div id="theater_nav">
            <a href="watchFilm">
                <div class="nav_noselected">
                    <div class="nav_mean">
                        <p>上映中</p>
                        <p>now playing</p>
                    </div>
                </div>
            </a>
            <a href="comingSoon">
                <div class="nav_selected">
                    <div class="nav_mean">
                        <p>公開予定</p>
                        <p>coming soon</p>
                    </div>
                </div>
            </a>
        </div>
    </section>
</template>

<script>
    export default {
        name: "tab"
    }
</script>
