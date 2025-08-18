<script setup lang="ts">
import type {Player} from "../types/player.ts";

const props = defineProps<{
  player: Player
}>()


</script>

<template>
  <div>
    игрок: {{ player.id }}
  </div>
</template>

<style scoped>

</style>