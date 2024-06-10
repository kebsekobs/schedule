import React from "react";
import { useForm } from "react-hook-form";
import { useAddTeacherMutation } from "../api/AddTeacherMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";

const AddTeacherModal = ({ isOpen, toggleModal }) => {
  const form = useForm();
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = form;

  const addTeacherMutation = useAddTeacherMutation();
  const onSubmit = (data) => {
    addTeacherMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  if (!isOpen) {
    return null;
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <div className={styles["modal-header"]}>
          <h2>Введите данные</h2>
          <Button onClick={toggleModal} styleFeature="close">
            <CloseSvg />
          </Button>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className={styles["groups-form"]}
        >
          <label>Введите ФИО</label>
          <input
            {...register("name", { required: true })}
            placeholder="Иванов И.И."
            className={styles["input"]}
          />
          {errors.inputId && (
            <span className={styles["error"]}>Это поле обязательно</span>
          )}
          <Button type="submit">Отправить</Button>
        </form>
      </div>
    </div>
  );
};

export default AddTeacherModal;
