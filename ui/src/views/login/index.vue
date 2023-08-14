<template>
  <el-form
  ref="loginForm"
  :model="loginForm"
  class="login-form"
  autocomplete="on"
  label-position="left"
  >
    <div class="title-container">
      <h3 class="title">
        Login 
      </h3>
    </div>

    <el-form-item prop="username">
      <span class="svg-container">
        <svg-icon name="user" />
      </span>
      <el-input
        ref="username"
        v-model="loginForm.username"
        name="username"
        type="text"
        autocomplete="on"
      />
      </el-form-item>
      <el-form-item prop="password">
        <span class="svg-container">
          <svg-icon name="password" />
        </span>
        <el-input
          :key="passwordType"
          ref="password"
          v-model="loginForm.password"
          :type="passwordType"
          name="password"
          autocomplete="on"
          @keyup.enter.native="handleLogin"
        />
        <span
          class="show-pwd"
          @click="showPwd"
        >
          <svg-icon
            :name="passwordType === 'password' ? 'eye-off' : 'eye-on'"
          />
        </span>
      </el-form-item>

      <el-button
        :loading="loading"
        type="primary"
        style="width: 100%; margin-bottom: 30px"
        @click.native.prevent="handleLogin"
      >
        Submit
      </el-button>
  </el-form>
</template>

<script src="./index.ts" lang="ts"></script>