<script>
import { useQuery } from "@vue/apollo-composable";
import gql from "graphql-tag";
import { watch } from "vue";

// export default {
//   setup() {
//     const { result, loading } = useQuery(gql`
//       query {
//         post(id: 1) {
//           id
//           title
//           body
//         }
//       }
//     `);

//     watch(result, (value) => {
//       console.log(value);
//     });

//     return {
//       result,
//       loading,
//     };
//   },
// };
</script>

<template>
  <!-- <div v-if="loading">Loading...</div>

  <ul v-else-if="result && result.post">
    <li>{{ result.post.title }}</li>
  </ul>  -->
  <h1>surprize mother</h1>
</template>
