<template>
    <section
        class="my-5 relative h-screen max-h-[500px] flex flex-col items-center justify-center text-center text-white py-0 px-3">
        <div class="video-docker absolute top-0 left-0 w-full h-full overflow-hidden">
            <video class="min-w-full min-h-full absolute object-cover" src="/hero.mp4" type="video/mp4" autoplay muted
                loop></video>
        </div>
        <div class="video-content space-y-2">

            <section class="">
                <div class="mx-auto  px-4  lg:flex  lg:items-center">
                    <div class="mx-auto  text-center">
                        <h1 class="text-3xl font-extrabold sm:text-5xl">
                            Global flavors, easy steps. <strong class="font-extrabold text-red-700 sm:block"> Elevate
                                your meals.</strong>
                        </h1>

                        <p class="mt-4 sm:text-xl/relaxed text-gray-300">
                            Turn your kitchen into a haven of deliciousness with easy-to-follow recipes for every
                            occasion.
                        </p>

                    </div>
                </div>
            </section>

        </div>
    </section>

</template>

<style>
.video-docker video {
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.video-docker::after {
    content: "";
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    background: rgba(0, 0, 0, 0.6);
    z-index: 1;
}

.video-content {
    z-index: 2;
}
</style>