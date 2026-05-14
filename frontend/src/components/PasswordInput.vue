<script setup lang="ts">
import { CFormInput, CFormLabel, CInputGroup, CInputGroupText } from '@coreui/vue'
import { ref } from 'vue'
const model = defineModel<string>()
const props = defineProps<{
  label?: string
  invalid?: boolean
  feedbackInvalid?: string
  class?: string
}>()
const show = ref<boolean>(false)
</script>

<template>
  <div :class="props.class">
    <CFormLabel v-if="props.label">{{ props.label }}</CFormLabel>
    <CInputGroup class="mb-2">
      <CInputGroupText style="cursor: pointer" @click="() => (show = !show)"
        ><span
          :class="{ pi: true, 'pi-eye': show, 'pi-eye-slash': !show }"
          style="font-size: 0.8rem"
        ></span
      ></CInputGroupText>
      <CFormInput
        :type="show ? 'text' : 'password'"
        v-model="model"
        :invalid="props.invalid"
        :feedback-invalid="props.feedbackInvalid"
      />
    </CInputGroup>
  </div>
</template>

<style scoped></style>
