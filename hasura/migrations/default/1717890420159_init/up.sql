SET check_function_bodies = false;
CREATE TABLE public."Comment" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    comment text NOT NULL,
    "recipeId" uuid NOT NULL,
    rating integer NOT NULL,
    "userId" uuid NOT NULL
);
CREATE TABLE public."Favorite" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "userId" uuid NOT NULL,
    "receipeId" uuid NOT NULL
);
CREATE TABLE public."Ingredients" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    "recipeId" uuid NOT NULL
);
CREATE TABLE public."Recipe" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    "userId" uuid NOT NULL,
    duration integer DEFAULT 0 NOT NULL
);
CREATE TABLE public."RecipeTag" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "recipeId" uuid NOT NULL,
    "tagId" uuid NOT NULL
);
CREATE TABLE public."Tags" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "tagName" text NOT NULL
);
CREATE TABLE public."recipeImages" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "imageLink" text NOT NULL,
    "recipeId" uuid NOT NULL
);
CREATE TABLE public."recipeSteps" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    description text NOT NULL,
    "order" integer NOT NULL,
    "recipeId" uuid NOT NULL
);
CREATE TABLE public."userLikes" (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "userId" uuid NOT NULL,
    "recipeId" uuid NOT NULL
);
CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    "firstName" text NOT NULL,
    "lastName" text NOT NULL,
    password text NOT NULL,
    "profileImg" text,
    "isPublic" boolean DEFAULT false NOT NULL,
    email text
);
ALTER TABLE ONLY public."Comment"
    ADD CONSTRAINT "Comment_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."Favorite"
    ADD CONSTRAINT "Favorite_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."Favorite"
    ADD CONSTRAINT "Favorite_receipeId_userId_key" UNIQUE ("receipeId", "userId");
ALTER TABLE ONLY public."Ingredients"
    ADD CONSTRAINT "Ingredients_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."RecipeTag"
    ADD CONSTRAINT "RecipeTag_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."Recipe"
    ADD CONSTRAINT "Recipe_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."Tags"
    ADD CONSTRAINT "Tags_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."recipeImages"
    ADD CONSTRAINT "recipeImages_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."recipeSteps"
    ADD CONSTRAINT "recipeSteps_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."userLikes"
    ADD CONSTRAINT "userLikes_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public."userLikes"
    ADD CONSTRAINT "userLikes_userId_recipeId_key" UNIQUE ("userId", "recipeId");
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public."Comment"
    ADD CONSTRAINT "Comment_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."Comment"
    ADD CONSTRAINT "Comment_userId_fkey" FOREIGN KEY ("userId") REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."Favorite"
    ADD CONSTRAINT "Favorite_receipeId_fkey" FOREIGN KEY ("receipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."Favorite"
    ADD CONSTRAINT "Favorite_userId_fkey" FOREIGN KEY ("userId") REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."Ingredients"
    ADD CONSTRAINT "Ingredients_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."RecipeTag"
    ADD CONSTRAINT "RecipeTag_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."RecipeTag"
    ADD CONSTRAINT "RecipeTag_tagId_fkey" FOREIGN KEY ("tagId") REFERENCES public."Tags"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."Recipe"
    ADD CONSTRAINT "Recipe_userId_fkey" FOREIGN KEY ("userId") REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."recipeImages"
    ADD CONSTRAINT "recipeImages_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."recipeSteps"
    ADD CONSTRAINT "recipeSteps_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."userLikes"
    ADD CONSTRAINT "userLikes_recipeId_fkey" FOREIGN KEY ("recipeId") REFERENCES public."Recipe"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE ONLY public."userLikes"
    ADD CONSTRAINT "userLikes_userId_fkey" FOREIGN KEY ("userId") REFERENCES public.users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
