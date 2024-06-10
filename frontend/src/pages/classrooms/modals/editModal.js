import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useClassroomByIdQuery } from "../api/getClassroomById";
import { useEditClassroomMutation } from "../api/editClassroomMutation";
import styles from "../../shared/style/modal.module.css";
import Button from "../../../components/button";
import {CloseSvg} from "../../../components/close-svg";

const EditGroupModal = ({ isOpen, toggleModal, id }) => {
  const classroomByIdQuery = useClassroomByIdQuery(id);
  const [isDataLoaded, setIsDataLoaded] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm({
    defaultValues: {
      classroomId: "",
      capacity: "",
    },
  });

  const editClassroomMutation = useEditClassroomMutation();
  const onSubmit = (data) => {
    editClassroomMutation.mutateAsync(data);
    toggleModal();
    reset();
  };

  useEffect(() => {
    if (classroomByIdQuery.data) {
      reset({
        id: classroomByIdQuery.data.id,
        classroomId: classroomByIdQuery.data.classroomId,
        capacity: classroomByIdQuery.data.capacity,
      });
      setIsDataLoaded(true);
    }
  }, [classroomByIdQuery.data, reset]);

  if (!isOpen) {
    return null;
  }

  if (!isDataLoaded) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <div className={styles["modal-header"]}>
          <h2>Измените данные</h2>
          <Button onClick={toggleModal} styleFeature="close">
            <CloseSvg />
          </Button>
        </div>
        <form
            onSubmit={handleSubmit(onSubmit)}
            className={styles["groups-form"]}
        >
          <label>Введите аудиторию</label>
          <input
              {...register("classroomId", { required: true })}
              placeholder="103(б)"
              className={styles["input"]}
          />
          {errors.classroomId && (
              <span className={styles["error"]}>Это поле обязательно</span>
          )}
          <label>Введите вместимость аудитории</label>
          <input
              type="number"
              {...register("capacity", { required: true, min: 1 })}
              placeholder="46"
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

export default EditGroupModal;
