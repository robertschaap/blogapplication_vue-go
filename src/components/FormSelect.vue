<template>
  <FormGroup :label="label">
    <select
      :class="$style.component"
      :value="value"
      @input="$emit('input', $event.target.value)">
      <option
        v-for="option in options"
        :key="option.value"
        :value="option.value">
        {{ option.label }}
      </option>
    </select>
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormSelect',
  props: {
    label: String,
    options: Array,
    value: String,
  },
  components: {
    FormGroup,
  },
};
</script>

<style lang="scss" module>
.component {
  display: block;
  width: 100%;
  background-color: white;
}
</style>
