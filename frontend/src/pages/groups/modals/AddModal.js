import React from "react";
import { useForm } from "react-hook-form";
import { useAddGroupsMutation } from "../api/AddGroupMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";

const AddGroupModal = ({data, isOpen, toggleModal }) => {
  const form = useForm();
  const existingGroupIds = data?.map(group => group.id) ?? []
  const {
    reset,
    register,
    handleSubmit,
    formState: { errors },
  } = form;

  const addGroupMutation = useAddGroupsMutation();
  const onSubmit = (data) => {
    addGroupMutation.mutateAsync(data);
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
            <label>Укажите степень</label>
            <label htmlFor="consent" className={styles["label"]}>
              <input
                  type="checkbox"
                  id="magistracy"
                  className={styles["checkbox"]}
                  {...register("magistracy")}
              />
              Магистратура
            </label>
            <label>Введите номер группы</label>
            <input
                type="number"
                {...register("groupId", { required: true })}
                placeholder="118"
                className={styles["input"]}
            />
            {errors.id && (
                <span className={styles["error"]}>Это поле обязательно</span>
            )}
            <label>Введите код группы</label>
            <input
                {...register("id", {
                  required: "Это поле обязательно",
                  validate: value => !existingGroupIds.includes(value) || "Такой ID уже существует"
                })}
                placeholder="БOЮ15-РПИ2101"
                className={styles["input"]}
            />
            {errors.id && (
                <span className={styles["error"]}>{errors.id.message}</span>
            )}
            <label>Введите количество студентов</label>
            <input
                type="number"
                {...register("capacity", { required: true, min: 1 })}
                placeholder="13"
                className={styles["input"]}
            />
            {errors.capacity && (
                <span className={styles["error"]}>
              {errors.capacity.type === "required"
                  ? "Это поле обязательно"
                  : "Значение должно быть больше 0"}
            </span>
            )}
            <Button type="submit">Отправить</Button>
          </form>
        </div>
      </div>
  );
};

export default AddGroupModal;
