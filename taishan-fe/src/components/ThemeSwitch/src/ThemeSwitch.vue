<script setup lang="ts">
import { useAppStore } from '@/store/modules/app'
import { ElSwitch } from 'element-plus'
import { useIcon } from '@/hooks/web/useIcon'
import { useDesign } from '@/hooks/web/useDesign'

const { getPrefixCls } = useDesign()

const emit = defineEmits(['change'])

const prefixCls = getPrefixCls('theme-switch')

const Sun = useIcon({ icon: 'emojione-monotone:sun', color: '#fde047' })

const CrescentMoon = useIcon({ icon: 'emojione-monotone:crescent-moon', color: '#fde047' })

const appStore = useAppStore()

// 初始化获取是否是暗黑主题
const isDark = false

// 设置switch的背景颜色
const blackColor = 'var(--el-color-black)'

const themeChange = (val: boolean) => {
  console.log(val)
  appStore.setIsDark(false)
  emit('change', false)
}
</script>

<template>
  <ElSwitch
    :class="prefixCls"
    v-model="isDark"
    inline-prompt
    :border-color="blackColor"
    :inactive-color="blackColor"
    :active-color="blackColor"
    :active-icon="Sun"
    :inactive-icon="CrescentMoon"
    @change="themeChange"
  />
</template>

<style lang="less" scoped>
:deep(.el-switch__core .el-switch__inner .is-icon) {
  overflow: visible;
}
</style>
