<template>
  <FormGroup :label="label">
    <select :class="$style.component">
      <option value="val">option</option>
    </select>
  </FormGroup>
</template>

<script>
import FormGroup from '@/components/FormGroup';

export default {
  name: 'FormSelect',
  props: {
    label: String,
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
